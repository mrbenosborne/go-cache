package main

import (
	"fmt"

	"github.com/mrbenosborne/go-cache"
)

func main() {

	// Create a new cache with a default
	// size of 0.
	c := cache.New(0)

	// Create a new store for flights.
	s := c.New("flights")

	// Add an item to the store with the key: LGW
	// and the value of: "London Gatwick".
	s.Add("LGW", "London Gatwick")

	// Retrieve the item from the store.
	airportName, _ := s.Get("LGW").String()
	fmt.Printf("The airport name is: %s\n", airportName)

	// Delete the item from the store.
	s.Delete("LGW")

}
