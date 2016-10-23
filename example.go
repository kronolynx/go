package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"./stampery"
)

func main() {
	events := stampery.Login("2d4cdee7-38b0-4a66-da87-c1ab05b43768")

	for event := range events {
		switch event.Type {
		case "ready":
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			digest := stampery.Hash("Hello blockchain!" + strconv.Itoa(r.Int()))
			stampery.Stamp(digest)
		case "proof":
			fmt.Println("\nProof")
			p := event.Data.(stampery.Proof)
			fmt.Printf("Version: %v\nSiblings: %v\nRoot: %v\n", p.Version, p.Siblings, p.Root)
			fmt.Printf("Anchor:\n  Chain: %v\n  Tx: %v\n", p.Anchor.Chain, p.Anchor.Tx)

		case "error":
			log.Fatalf("%v\n", event.Data)
		}
	}
}
