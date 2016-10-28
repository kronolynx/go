# Stampery
Stampery API for go. Notarize all your data using the blockchain!

## Installation

Add this line to your imports:

```go
"github.com/stampery/go/stampery"
```

And then execute:

    $ go get


## Usage
```go
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
    // Sign up and get your secret token at https://api-dashboard.stampery.com
	events := stampery.Login("user-secret")

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
```

# Official implementations
- [NodeJS](https://github.com/stampery/node)
- [PHP](https://github.com/stampery/php)
- [ruby](https://github.com/stampery/ruby)
- [Python](https://github.com/stampery/python)
- [Elixir](https://github.com/stampery/elixir)
- [Java](https://github.com/stampery/java)
- [Go](https://github.com/stampery/go)

# Feedback

Ping us at support@stampery.com and we’ll help you! 😃


## License

Code released under
[the MIT license](https://github.com/stampery/js/blob/master/LICENSE).

Copyright 2016 Stampery
