package avl

import (
	linkedlist "github.com/salmanrf/go-dsa/linked-list"
)

type CompareFunc[T any] func (a, b T) int

type AVLNode[T any] struct {
	Data 				T
	Left 				*AVLNode[T]
	Right 			*AVLNode[T]
	Height 			int
	compareFunc CompareFunc[T]
}

func NewAVL[T any](data T, cf CompareFunc[T]) *AVLNode[T] {
	node := &AVLNode[T]{
		Data: data,
		Left: nil,
		Right: nil,
		Height: 1,
		compareFunc: cf,
	}

	return node
}

func (r *AVLNode[T]) Search(val T) *AVLNode[T] {
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

func (r *AVLNode[T]) CollectInorder(collection linkedlist.SingleLinkedList[*AVLNode[T]]) linkedlist.SingleLinkedList[*AVLNode[T]] {
	if r.Left != nil {
		collection = r.Left.CollectInorder(collection)
	}

	if collection == nil {
		collection = linkedlist.New(r)
	} else {
		collection.Append(r)
	}

	if r.Right != nil {
		collection = r.Right.CollectInorder(collection)
	}

	return collection
}

func (r *AVLNode[T]) CalculateHeight() int {
	if r == nil {
		return 0
	}

	lh := 0
	rh := 0

	if r.Left != nil {
		lh = r.Left.Height
	}

	if r.Right != nil {
		rh = r.Right.Height
	}
	
	return 1 + max(lh, rh)
}

func rebalanceTree[T any](root *AVLNode[T]) *AVLNode[T] {
	if root == nil {
		return nil
	}
	
	root.Height = root.CalculateHeight()
	balanceFactor := root.GetBalanceFactor()

	// * Left Imbalance
	if balanceFactor > 1 {
		lbf := 0
		
		if root.Left != nil {
			lbf = root.Left.GetBalanceFactor()
		} 
			
		// * Left left imbalance (Heavier on left subtree)
		// * The BF will be positive
		if lbf > 0 {
			return llRotate(root)
		}

		// * Left right imbalance (Heavier on left subtree's right)
		// * The BF will be negative
		if lbf < 0 {
			return lrRotate(root)
		}
	}

	// * Right Imbalance
	if balanceFactor < -1 {
		rbf := 0

		if root.Right != nil {
			rbf = root.Right.GetBalanceFactor()
		}

		// * Right right imbalance
		if rbf < 0 {
			return rrRotate(root)
		}

		// * Right left imbalance
		if rbf > 0 {
			return rlRotate(root)
		}
	}

	return root
}

func (r *AVLNode[T]) Insert(val T) *AVLNode[T] {
	if r.compareFunc(r.Data, val) < 0 {
		if r.Left == nil {
			r.Left = NewAVL(val, r.compareFunc)
		} else {
			r.Left = r.Left.Insert(val)
		}
	}

	if r.compareFunc(r.Data, val) > 0 {
		if r.Right == nil {
			r.Right = NewAVL(val, r.compareFunc)
		} else {
			r.Right = r.Right.Insert(val)
		}
	}
	
	return rebalanceTree(r)
}

func (r *AVLNode[T]) Delete(val T) *AVLNode[T] {
	compareResult := r.compareFunc(r.Data, val)
	
	if compareResult < 0 {
		if r.Left != nil {
			r.Left = r.Left.Delete(val)

			return rebalanceTree(r)
		}
	}

	if compareResult > 0 {
		if r.Right != nil {
			r.Right = r.Right.Delete(val)
			
			return rebalanceTree(r)
		}
	}

	if r.Left == nil {
		return r.Right
	}

	if r.Right == nil {
		return r.Left
	}
	
	predecessor := r.InorderPredecessor()

	if predecessor != nil {
		r.Data = predecessor.Data
		r.Left = r.Left.Delete(predecessor.Data)

		return rebalanceTree(r)
	}

	successor := r.InorderSuccessor()

	if successor != nil {
		r.Data = successor.Data
		r.Right = r.Right.Delete(successor.Data)

		return rebalanceTree(r)
	}

	return nil
}

func (r *AVLNode[T]) GetBalanceFactor() int {
	bf := 0

	lh := 0
	rh := 0

	if r.Left != nil {
		lh = r.Left.CalculateHeight()
	}

	if r.Right != nil {
		rh = r.Right.CalculateHeight()
	}

	bf = lh - rh

	return bf
}

func llRotate[T any](r *AVLNode[T]) *AVLNode[T] {
	rl := r.Left
	rlr := rl.Right 

	rl.Right = r
	r.Left = rlr

	r.Height = r.CalculateHeight()
	rl.Height = rl.CalculateHeight()

	return rl
}

func lrRotate[T any](p *AVLNode[T]) *AVLNode[T] {
	pl := p.Left
	plr := pl.Right

	p.Left = plr.Right
	pl.Right = plr.Left

	plr.Right = p
	plr.Left = pl

	p.Height = p.CalculateHeight()
	pl.Height = pl.CalculateHeight()
	plr.Height = plr.CalculateHeight()

	return plr
}

func rrRotate[T any](p *AVLNode[T]) *AVLNode[T] {
	pr := p.Right
	prl := p.Right.Left

	pr.Left = p
	p.Right = prl

	p.Height = p.CalculateHeight()
	pr.Height = pr.CalculateHeight()

	return pr
}

func rlRotate[T any](p *AVLNode[T]) *AVLNode[T] {
	pr := p.Right
	prl := p.Right.Left

	p.Right = prl.Left
	pr.Left = prl.Right

	prl.Left = p
	prl.Right = pr

	p.Height = p.CalculateHeight()
	pr.Height = pr.CalculateHeight()
	prl.Height = prl.CalculateHeight()
	
	return prl
}

func (r *AVLNode[T]) InorderSuccessor() *AVLNode[T] {
	if r.Right == nil {
		return nil
	}

	return r.Right.Min()
}

func (r *AVLNode[T]) InorderPredecessor() *AVLNode[T] {
	if r.Left == nil {
		return nil
	}

	return r.Left.Max()
}

func (r *AVLNode[T]) Min() *AVLNode[T] {
	if r == nil {
		return nil
	}
	
	if r.Left == nil {
		return r
	}

	return r.Left.Min()
}

func (r *AVLNode[T]) Max() *AVLNode[T] {
	if r == nil {
		return nil
	}
	
	if r.Right == nil {
		return r
	}

	return r.Right.Max()
}

func (r *AVLNode[T]) GetHeight() int {
	if r == nil {
		return 0
	}

	leftH := 0
	rightH := 0

	if r.Left != nil {
		leftH = r.Left.GetHeight()
	}
	
	if r.Right != nil {
		rightH = r.Right.GetHeight()
	}

	return 1 + max(leftH, rightH)
}