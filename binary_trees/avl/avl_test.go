package avl

import (
	"fmt"
	"testing"

	linkedlist "github.com/salmanrf/go-dsa/linked-list"
)

func compareFuncInt(a, b int) int {
	return b - a
}

func TestNewAVL(t *testing.T) {
	want := AVLNode[int]{
		Data: 10,
		Left: nil,
		Right: nil,
		Height: 1,
	}

	got := NewAVL(10, compareFuncInt)

	if got.Data != want.Data {
		t.Errorf("Got %v, want %v", got.Data, want.Data)
	}

	if got.Left != want.Left {
		t.Errorf("Got %v, want %v", got.Left, want.Left)
	}

	if got.Right != want.Right {
		t.Errorf("Got %v, want %v", got.Right, want.Right)
	}

	if got.Height != want.Height {
		t.Errorf("Got %v, want %v", got.Height, want.Height)
	}
}

func TestAVLInsert(t *testing.T) {
	root := NewAVL(10, compareFuncInt)

	root = root.Insert(8)
	root = root.Insert(12)
	root = root.Insert(14)
	root = root.Insert(13)

	list := root.CollectInorder(nil)

	i := 0
	
	list.Traverse(func (n *linkedlist.ListNode[*AVLNode[int]]) {
		elements := []int{8, 10, 12, 13, 14}
		t.Run("Element should exist and in order", func(t *testing.T) {
			want := elements[i]

			got := n.Val
				
			if got == nil  {
				t.Errorf("got %v, want %v", got, want)
			}
		
			if got.Data != want  {
				t.Errorf("got %v, want %v", got.Data, want)
			}

			i++
		})
	})
}


func TestInsertHeight(t *testing.T) {
	root := NewAVL(100, compareFuncInt)

	t.Run("New tree should have height 1", func (t *testing.T) {
		want := 1
		got := root.Height

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	root = root.Insert(20)
	root = root.Insert(200)

	t.Run("Level 2", func (t *testing.T) {
		want := 2
		got := root.Height

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	root = root.Insert(10)

	t.Run("Level 3", func (t *testing.T) {
		want := 3
		got := root.Height

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	root = root.Insert(25)

	t.Run("Still level 3", func (t *testing.T) {
		want := 3
		got := root.Height

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestInsertLeftLeftImbalance(t *testing.T) {
	root := NewAVL(100, compareFuncInt)

	root = root.Insert(50)
	root = root.Insert(24)
	root = root.Insert(12)
	root = root.Insert(6)
	root = root.Insert(2)
	root = root.Insert(0)

	t.Run("Tree is rebalanced after LL rotation", func (t *testing.T) {
		want := 12
		got := root.Data

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Elements are still inorder", func (t *testing.T) {
		elements := []int{0, 2, 6, 12, 24, 50, 100}

		list := root.CollectInorder(nil)

		i := 0

		list.Traverse(func (n *linkedlist.ListNode[*AVLNode[int]]) {
			want := elements[i]

			got := n.Val.Data

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
			
			i++
		})
	})
}

func TestInsertLeftRightImbalance(t *testing.T) {
	root := NewAVL(50, compareFuncInt)

	root = root.Insert(40)
	root = root.Insert(45)

	want := 45
	got := root.Data

	if got != want {
		t.Errorf("got %d, want %d", got, want)	
	}
}

func TestInsertRightRightImbalance(t *testing.T) {
	root := NewAVL(10, compareFuncInt)

	root = root.Insert(5)
	root = root.Insert(20)
	root = root.Insert(15)
	root = root.Insert(30)
	root = root.Insert(40)

	t.Run("Tree is RR rotated", func (t *testing.T) {
		want := 20
		got := root.Data

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Ensure subtree is assigned to correct position", func(t *testing.T) {
		want := 15
		got := root.Left.Right.Data

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestInsertRightLeftImbalance(t *testing.T) {
	root := NewAVL(50, compareFuncInt)

	root = root.Insert(40)
	root = root.Insert(60)
	root = root.Insert(70)
	// * Imbalance
	root = root.Insert(65)

	t.Run("Tree is RL Rotated", func (t *testing.T) {
		want := 65
		got := root.Right.Data

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func TestDelete(t *testing.T) {
	root := NewAVL(50, compareFuncInt)

	root = root.Insert(40)
	root = root.Insert(60)
	root = root.Insert(70)
	root = root.Insert(65)

	root = root.Delete(60)

	t.Run("Deleted node is gone", func (t *testing.T) {
		var want *AVLNode[int] = nil
		got := root.Search(60)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Root hasn't changed", func (t *testing.T) {
		want := 50
		got := root.Data

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	
	root = root.Delete(50)

	t.Run("Deleted root is replaced by predecessor first", func (t *testing.T) {
		want := 65
		got := root.Data

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("Tree is rebalanced after deletion (LL)", func (t *testing.T) {
		root := NewAVL(50, compareFuncInt)

		root = root.Insert(25)
		root = root.Insert(100)
		root = root.Insert(200)

		root = root.Delete(25)

		want := 100
		got := root.Data

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

		list := root.CollectInorder(nil)

		elements := []int{50, 100, 200}
		i := 0
		
		list.Traverse(func (lnode *linkedlist.ListNode[*AVLNode[int]]) {
			t.Run(fmt.Sprintf("Elements are in order %d\n", i), func (t *testing.T) {
				want := elements[i]
	
				got := lnode.Val

				if got.Data != want {
					t.Errorf("got %d, want %d", got.Data, want)
				}

				gotBf := got.GetBalanceFactor() 

				if gotBf < -1 || gotBf > 1 {
					t.Errorf("got balance factor %d, want [-1, 1]", gotBf)
				}
	
				i++
			})
		})
	})

	t.Run("Tree is rebalanced after deletion (LR)", func (t *testing.T) {
		root := NewAVL(50, compareFuncInt)

		root = root.Insert(25)
		root = root.Insert(100)
		root = root.Insert(200)
		root = root.Insert(40)

		root = root.Delete(100)
		root = root.Delete(200)

		want := 40
		got := root.Data

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}

		list := root.CollectInorder(nil)

		elements := []int{25, 40, 50}
		i := 0
		
		list.Traverse(func (lnode *linkedlist.ListNode[*AVLNode[int]]) {
			t.Run(fmt.Sprintf("Elements are in order %d\n", i), func (t *testing.T) {
				want := elements[i]
	
				got := lnode.Val

				if got.Data != want {
					t.Errorf("got %d, want %d", got.Data, want)
				}

				gotBf := got.GetBalanceFactor() 

				if gotBf < -1 || gotBf > 1 {
					t.Errorf("got balance factor %d, want [-1, 1]", gotBf)
				}
	
				i++
			})
		})
	})
}