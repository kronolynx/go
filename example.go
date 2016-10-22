package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"./stampery"
)

func main() {
	events := stampery.Login("a0ad0ee3-2466-43db-9b88-5185bd2cc40b")

	for event := range events {
		switch event.Type {
		case "ready":
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			digest := stampery.Hash("Hello blockchain!" + strconv.Itoa(r.Int()))
			stampery.Stamp(digest)
		case "proof":
			fmt.Println(event.Data)
		case "error":
			fmt.Println(event.Data)
		}
	}
}
