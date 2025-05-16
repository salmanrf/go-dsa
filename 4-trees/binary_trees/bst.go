package binary_trees

type CompareFunc[T any] func (a, b T) int

type BSTNode[T any] struct {
	Data T
	Left *BSTNode[T]
	Right *BSTNode[T]
	compareFunc CompareFunc[T]
}

func New[T any](data T, cf CompareFunc[T]) *BSTNode[T] {
	node := &BSTNode[T]{
		Data: data,
		Left: nil,
		Right: nil,
		compareFunc: cf,
	}

	return node
}

func (r *BSTNode[T]) Insert(val T) *BSTNode[T] {
	if r.compareFunc(val, r.Data) < 0 {
		if r.Left == nil {
			r.Left = New(val, r.compareFunc)
		}
		
		r.Left = r.Left.Insert(val)
	}

	if r.compareFunc(val, r.Data) > 0 {
		if r.Right == nil {
			r.Right = New(val, r.compareFunc)
		}
		
		r.Right = r.Right.Insert(val)
	}

	return r
}

func (r *BSTNode[T]) Search(val T) *BSTNode[T] {
	if r.compareFunc(val, r.Data) == 0 {
		return r
	} 

	if r.compareFunc(val, r.Data) < 0 {
		if r.Left == nil {
			return nil
		}
		
		return r.Left.Search(val)
	}

	if r.compareFunc(val, r.Data) > 0 {
		if r.Right == nil {
			return nil
		}
		
		return r.Right.Search(val)
	}

	return nil
}

func (r *BSTNode[T]) Min() *BSTNode[T] {
	if r == nil {
		return nil
	}
	
	if r.Left == nil {
		return r
	}

	return r.Left.Min()
}

func (r *BSTNode[T]) Max() *BSTNode[T] {
	if r == nil {
		return nil
	}
	
	if r.Right == nil {
		return r
	}

	return r.Right.Max()
}

// TODO: Use parent as successor
func (r *BSTNode[T]) InorderSuccessor() *BSTNode[T] {
	if r.Right == nil {
		return nil
	}

	return r.Right.Min()
}

func (r *BSTNode[T]) InorderPredecessor() *BSTNode[T] {
	if r.Left == nil {
		return nil
	}

	return r.Left.Max()
}

func (r *BSTNode[T]) Delete(val T) *BSTNode[T] {
	if r.compareFunc(val, r.Data) == 0 {
		if r.Left == nil && r.Right == nil {
			return nil
		}

		if r.Left != nil && r.Right == nil {
			return r.Left
		}

		if r.Right != nil && r.Left == nil {
			return r.Right
		}

		// * Has both children
		// * Try to find inorder successor first
		inSuccessor := r.InorderSuccessor()

		if inSuccessor != nil {
			r.Right = r.Delete(inSuccessor.Data)
			r.Data = inSuccessor.Data

			return r
		}

		inPredecessor := r.InorderPredecessor()

		if inPredecessor != nil {
			r.Left = r.Delete(inPredecessor.Data)
			r.Data = inPredecessor.Data

			return r
		}
	}

	if r.compareFunc(val, r.Data) < 0 {
		if r.Left != nil {
			r.Left = r.Left.Delete(val)
		}		
	}

	if r.compareFunc(val, r.Data) > 0 {
		if r.Right != nil {
			r.Right = r.Right.Delete(val)
		}
	}

	return r
} 