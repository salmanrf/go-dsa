package binary_trees

type TraversalCompare[T any] func (new T, data *Node[T]) int

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
	if root == nil {
		return &Node[T]{
			Data: data,
		};
	}

	if compare(data, root) <= 0 {
		root.Left = BSTInsert(root.Left, data, compare);
	}

	if compare(data, root) > 0 {
		root.Right = BSTInsert(root.Right, data, compare);
	}

	return root;
}

func GetInorderSuccessor[T any](root *Node[T], value T, parentCompare TraversalCompare[T], compare TraversalCompare[T]) *Node[T] {
	node := BSTSearch(
		root,
		value,
		compare,
	)

	if node == nil {
		return nil
	}

	if node.Right == nil {
		
	}
	
	// ? Get parent of target node, in case it doesn't have a right sub-tree
	parent := BSTSearch(
		root,
		value,
		parentCompare,
	)
}


