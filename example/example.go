package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/stampery/go/stampery"
)

func main() {
	// Login to the stampery API
	events := stampery.Login("2d4cdee7-38b0-4a66-da87-c1ab05b43768")

	for event := range events {
		switch event.Type {
		case "ready":
			// In this case we are going to add a random number to the string
			// to generate a different hash each time.
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			digest := stampery.Hash("Hello blockchain!" + strconv.Itoa(r.Int()))
			// Stamp the hash
			stampery.Stamp(digest)

		case "proof":
			fmt.Println("\nProof")
			p := event.Data.(stampery.Proof)
			fmt.Println("Hash: ", p.Hash)
			fmt.Printf("Version: %v\nSiblings: %v\nRoot: %v\n", p.Version, p.Siblings, p.Root)
			fmt.Printf("Anchor:\n  Chain: %v\n  Tx: %v\n", p.Anchor.Chain, p.Anchor.Tx)

		case "error":
			log.Fatalf("%v\n", event.Data)
		}
	}
}
