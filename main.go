package main

import (
	"fmt"

	"github.com/salmanrf/go-dsa/binary_trees/avl"
)

func compareFunc(a, b int) int {
	return b - a
}

func main() {
	root := avl.NewAVL(10, compareFunc)

	root = root.Insert(40)
	root = root.Insert(60)
	root = root.Insert(70)
	root = root.Insert(65)

	root = root.Delete(60)

	fmt.Println("ROOT", root.Data)
}