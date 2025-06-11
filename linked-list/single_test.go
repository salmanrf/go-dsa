package linkedlist

import (
	"fmt"
	"testing"
)

func compareFuncInt(a, b int) int {
	return b - a
}

func TestNew(t *testing.T) {
	t.Run("Should return the head of the list", func (t *testing.T) {
		want := 10
		got := New(10)
		
		if got == nil {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Should return the head from calling Get", func (t *testing.T) {
		wantExists, wantValue := true, 10
		
		head := New(10)

		gotExists, gotNode := head.Get(0)
		
		if gotExists != wantExists {
			t.Errorf("got %v, want %v", gotExists, wantExists)
		}

		if gotNode.Val != wantValue {
			t.Errorf("got %v, want %v", gotNode.Val, wantValue)
		}
	})
}

func TestGet(t *testing.T) {
	elements := []int{1, 2, 3, 4, 5, 6}

	tests := []struct{
		index int
		val		int
	}{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},		
	}

	head := New(elements[0])

	for _, el := range(elements[1:]) {
		head.Append(el)
	}

	for _, tt := range(tests) {
		t.Run(fmt.Sprintf("Index %d, val %d", tt.index, tt.val), func (t *testing.T) {
			wantExists, wantVal := true, tt.val

			gotExists, gotNode := head.Get(tt.index)

			if gotExists != wantExists {
				t.Errorf("got %v, want %v", gotExists, wantExists)
			}

			if gotNode.Val != wantVal {
				t.Errorf("got %v, want %v", gotNode.Val, wantVal)
			}
		})
	}
}

func TestFind(t *testing.T) {
	list := New(0)

	t.Run("Should return found node", func (t *testing.T) {
		got := list.Find(0, func (a, b int) bool { return a == b })
		want := 0

		if got == nil {
			t.Errorf("got %v, want non nil", got)
		}
		
		if got.Val != want {
			t.Errorf("got %v, want %v", got.Val, want)
		}
	})

	list.Append(100)
	list.Append(200)
	list.Append(-100)

	t.Run("Should return found node", func (t *testing.T) {
		got := list.Find(200, func (a, b int) bool { return a == b })
		want := 200

		if got.Val != want {
			t.Errorf("got %v, want %v", got.Val, want)
		}
	})

	t.Run("Should return nil when not found", func (t *testing.T) {
		got := list.Find(99, func (a, b int) bool { return a == b })
		var want *ListNode[int] = nil

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestAppendSorted(t *testing.T) {
	list := New(0)

	list.AppendSorted(-100, compareFuncInt)
	list.AppendSorted(100, compareFuncInt)
	list.AppendSorted(20, compareFuncInt)
	list.AppendSorted(10, compareFuncInt)
	list.AppendSorted(200, compareFuncInt)

	t.Run("Should append elements in ASC order", func (t *testing.T) {
		sorted := []int{-100, 0, 10, 20, 100, 200}

		i := 0
		
		list.Traverse(func (nd *ListNode[int]) {
			got := nd.Val
			want := sorted[i]

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}

			i++
		})
	})
}