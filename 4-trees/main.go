package main

import (
	"fmt"

	"github.com/salmanrf/trees/binary_trees"
)

func main() {
	// myarr := []int{1, 2, 3, 5, 6, 7, 8, 9}
	
	// mytree := binary_trees.CompleteBTFromArray(myarr);

	// fmt.Println("Root: ", myarr[0])
	// rootsLeft := (2 * 0) + 1
	// rootsRight := (2 * 0) + 2
	// rootsRightsParent := (rootsRight - 1) / 2

	// fmt.Println("Root's left: ", myarr[rootsLeft])
	// fmt.Println("Root's right", myarr[rootsRight])
	// fmt.Println("Root's right's parent", myarr[rootsRightsParent])
	
	// last := len(myarr) - 1
	// lastsParent := (last - 1) / 2

	// fmt.Println("Last", myarr[last])
	// fmt.Println("Last's parent", myarr[lastsParent])

	// binary_trees.TraversePreOrder(mytree, func (n *binary_trees.Node[int]) {
	// 	fmt.Println(n.Data)
	// })

	myarr2 := []int{8, 3, 10, 1, 6, 14, 4, 7};

	root2 := binary_trees.BSTFromArray(myarr2, func (a, b int) int {
		return a - b
	})

	fmt.Println("Root", root2.Data);
	fmt.Println("Left", root2.Left.Data);
	fmt.Println("Right", root2.Right.Data);
	fmt.Println("Left's Left", root2.Left.Left.Data);
	fmt.Println("Left's Right", root2.Left.Right.Data);
}