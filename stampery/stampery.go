package stampery

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/msgpack-rpc/msgpack-rpc-go/rpc"
	"github.com/streadway/amqp"
	"github.com/ugorji/go/codec"
	"golang.org/x/crypto/sha3"
)

var apiClient *rpc.Session
var events = make(chan Event)

//Login to the Stampery API
func Login(params ...string) chan Event {
	branch := "prod"
	if len(params) == 2 {
		branch = params[1]
	}

	clientID := getClientID(params[0])
	go apiLogin(clientID, params[0], branch)
	go amqpLogin(clientID, branch)
	return events
}

// Hash a string with sha3 512
func Hash(data string) string {
	hash := sha3.Sum512([]byte(data))
	return hex.EncodeToString(hash[:])
}

// Stamp a string
func Stamp(data string) {
	fmt.Printf("Stamping \n%v\n", data)
	_, err := apiClient.Send("stamp", strings.ToUpper(data))
	if err != nil {
		fmt.Println("error", err)
	}
}

func apiLogin(clientID, secret, branch string) {
	var address string
	if branch == "prod" {
		address = "api.stampery.com:4000"
	} else {
		address = "api-beta.stampery.com:4000"
	}
	conn, err := net.Dial("tcp", address)
	failOnError(err, "Couldn't connect to Api!")
	apiClient = rpc.NewSession(conn, true)
	isLogged, xerr := apiClient.Send("stampery.3.auth", clientID, secret)
	failOnError(xerr, "Couldn't connect to Api")
	if isLogged.Bool() {
		fmt.Println("logged ", clientID)
	}
}

func amqpLogin(clientID, branch string) {
	var url string
	if branch == "prod" {
		url = "amqp://consumer:9FBln3UxOgwgLZtYvResNXE7@young-squirrel.rmq.cloudamqp.com/ukgmnhoi"
	} else {
		url = "amqp://consumer:9FBln3UxOgwgLZtYvResNXE7@young-squirrel.rmq.cloudamqp.com/beta"
	}

	conn, err := amqp.Dial(url)
	failOnError(err, "Couldn't connect to Rabbit!")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open Rabbit channel!")

	fmt.Println("[QUEUE] Connected to Rabbit!")
	events <- Event{"ready", nil}
	msgs, err := ch.Consume(clientID+"-clnt", clientID+"-clnt", true, false, false, false, nil)
	failOnError(err, "Failed to register consumer!")

	for d := range msgs {
		var h codec.MsgpackHandle

		dec := codec.NewDecoderBytes(d.Body, &h)

		var v Proof
		v.Hash = d.RoutingKey
		// if the Proof doesn't contain siblings the conversion would fail
		// then we need to use a temporary struct to handle the error
		if err := dec.Decode(&v); err != nil {
			dec.ResetBytes(d.Body)
			var temp temp
			if err := dec.Decode(&temp); err != nil {
				failOnError(err, "Couldn't decode the proof :(")
			} else {
				v.Version = temp.Version
				v.Root = temp.Root
				v.Anchor = temp.Anchor
				events <- Event{"proof", v}
			}
		} else {
			events <- Event{"proof", v}
		}
	}
}

func getClientID(secret string) string {
	hasher := md5.New()
	hasher.Write([]byte(secret))
	return hex.EncodeToString(hasher.Sum(nil))[:15]
}

func failOnError(err error, msg string) {
	if err != nil {
		events <- Event{"error", err.Error() + " " + msg}
		log.Fatalf("%s %s", err, msg)
	}
}
