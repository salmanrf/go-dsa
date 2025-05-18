package binary_trees

import (
	"fmt"

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

func (r *AVLNode[T]) CollectInorder(collection linkedlist.SingleLinkedList[T]) linkedlist.SingleLinkedList[T] {
	if r.Left != nil {
		collection = r.Left.CollectInorder(collection)
	}

	if collection == nil {
		collection = linkedlist.New(r.Data)
	} else {
		collection.Append(r.Data)
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

	// * Maintain tree balance
	r.Height = r.CalculateHeight()
	balanceFactor := r.GetBalanceFactor()

	// * Left Imbalance
	if balanceFactor > 1 {
		lbf := 0
		
		if r.Left != nil {
			lbf = r.Left.GetBalanceFactor()
		} 
			
		// * Left left imbalance (Heavier on left subtree)
		// * The BF will be positive
		if lbf > 0 {
			return llRotate(r)
		}

		// * Left right imbalance (Heavier on left subtree's right)
		// * The BF will be negative
		if lbf < 0 {
			return lrRotate(r)
		}
	}

	// * Right Imbalance
	if balanceFactor < -1 {
		rbf := 0

		if r.Right != nil {
			rbf = r.Right.GetBalanceFactor()
		}

		// * Right right imbalance
		if rbf < 0 {
			return rrRotate(r)
		}

		// * Right left imbalance
		if rbf > 0 {
			return rlRotate(r)
		}
	}

	return r
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
	fmt.Println(p.Data)
	
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