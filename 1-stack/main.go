package main

import (
	"fmt"

	"github.com/salmanrf/data-structures/1-stack/stack"
)

func main() {
	my_s := stack.New[string](5)

	fmt.Println(my_s.Peek())
	
	my_s.Push("Amazing!!!")

	fmt.Println(my_s.Peek())
	fmt.Println(my_s.Peek())

	my_s.Push("Stack is easy")

	fmt.Println(my_s.Peek())

	my_s.Pop()

	fmt.Println(my_s.Peek())
	fmt.Println(my_s.Peek())

	fmt.Println("==========================================")

	my_s = stack.New[string](3)
	my_s.Push("Yay")
	my_s.Push("Woah")
	my_s.Push("Oww")
	my_s.Push("Okay")

	fmt.Println(my_s.Peek())
}