package main

import (
	"fmt"

	"github.com/salmanrf/linked_lists/single"
)

func main() {
	ll := single.New[int](1)

	ll.Append(2)	
	ll.Append(99)
	ll.Append(-1000)

	fmt.Printf("Value at %d: ", 0)
	fmt.Println(ll.Get(0))

	fmt.Printf("Value at %d: ", 1)
	fmt.Println(ll.Get(1))

	fmt.Printf("Value at %d: ", 2)
	fmt.Println(ll.Get(2))
	
	fmt.Printf("Value at %d: ", 3)
	fmt.Println(ll.Get(3))

	fmt.Printf("Value at %d: ", -5)
	fmt.Println(ll.Get(-5))
}
