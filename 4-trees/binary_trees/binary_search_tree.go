package binary_trees

type TraversalCompare[T any] func (new T, node *Node[T]) int

func BSTFromArray[T any](items []T, compare TraversalCompare[T]) *Node[T] {
	size := len(items);

	var root *Node[T] = nil;

	for i := 0; i < size; i++ {
		root = BSTInsert(root, items[i], compare);
	}

	return root;
}

func BSTSearch[T any](root *Node[T], target T, compare TraversalCompare[T]) *Node[T] {
	if root == nil {
		return nil;
	}

	if compare(target, root) == 0 {
		return root;
	}

	if compare(target, root) < 0 {
		return BSTSearch(root.Left, target, compare);
	}

	if compare(target, root) > 0 {
		return BSTSearch(root.Right, target, compare);
	}

	return nil;
}

func BSTInsert[T any](root *Node[T], data T, compare TraversalCompare[T]) *Node[T] {
	// * 1. If root is nil, create a new node (root)
	// * 2. If data is less than root.data, insert to left,
	// * 3. If data is greater than root.data, insert to right,
	// Can we omit returning root at the end ? 
	// This should have worked because we're working with pointers. If we find the correct node to insert into, 
	// it shouldn't matter because the original pointers (root.left or root.right) aren't updated directly. 
	// Essentially, we are performing a search for the correct node, and at the end, we create a new node and return it.
	// Ahhh, I get it now! We need to create a new node and return it at some point, and this return has to be accepted 
	// somewhere, either as root.left or root.right. Once we finish at the current frame, we still need to move back up.
	// If we don't return the root (current node), the calling frame will receive nil, which would destroy the hierarchy.

	
	if root == nil {
		return &Node[T]{
			Data: data,
		}
	}

	if(compare(data, root)) == 0 {
		return root
	} 

	if compare(data, root) <= 0 {
		root.Left = BSTInsert(root.Left, data, compare)
	} else {
		root.Right = BSTInsert(root.Right, data, compare)
	}

	return root
}




