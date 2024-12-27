package binary_trees

type Node[T any] struct {
	Data T
	Left *Node[T]
	Right *Node[T]
}

// ! INCORRECT
func FromArray[T any](arr []T) *Node[T] {
	length := len(arr)

	if length == 0 {
		return nil
	}

	root := &Node[T]{
		Data: arr[0],
		Left: nil,
		Right: nil,
	}

	isLeft := true
	current := root

	for i := 1; i < len(arr); i++ {
		if isLeft {
			current.Left = &Node[T]{
				Data: arr[i],
			}

			isLeft = false
		} else {
			current.Right = &Node[T]{
				Data: arr[i],
			}

			current = current.Left

			isLeft = true
		}
	}

	return root
}

func TraversePreOrder[T any](root *Node[T], fn func (n *Node[T])) {
	if root != nil {
		fn(root)
		TraversePreOrder(root.Left, fn)		
		TraversePreOrder(root.Right, fn)
	}
}
