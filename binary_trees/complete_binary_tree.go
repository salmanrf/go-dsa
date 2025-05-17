package binary_trees

func CompleteBTFromArray[T any](arr []T) *Node[T] {
	root := &Node[T]{
		Data: arr[0],
		Left: nil,
		Right: nil,
	}

	length := len(arr);

	createChildren(arr, root, 0, length);

	return root;
} 

func createChildren[T any](arr []T, root *Node[T], index, length int) {
		if index < length {
			leftIndex := 2 * index + 1;
			rightIndex := 2 * index + 2;
	
			if leftIndex < length {
				root.Left = &Node[T]{
					Data: arr[leftIndex],
				}
			}

			if rightIndex < length {
				root.Right = &Node[T]{
					Data: arr[rightIndex],
				}
			}

			createChildren(arr, root.Left, leftIndex, length);
			createChildren(arr, root.Right, rightIndex, length);
		}
	}

func TraversePreOrder[T any](root *Node[T], fn func (n *Node[T])) {
	if root != nil {
		fn(root)
		TraversePreOrder(root.Left, fn)		
		TraversePreOrder(root.Right, fn)
	}
}

func TraverseInOrder[T any](root *Node[T], fn func (n *Node[T])) {
	if root != nil {
		TraverseInOrder(root.Left, fn)		
		fn(root)
		TraverseInOrder(root.Right, fn)
	}
}

func TraversePostOrder[T any](root *Node[T], fn func (n *Node[T])) {
	if root != nil {
		TraversePostOrder(root.Left, fn)		
		TraversePostOrder(root.Right, fn)
		fn(root)
	}
}