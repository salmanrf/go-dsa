package linkedlist

import (
	"fmt"
	"testing"
)

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

		if gotNode.val != wantValue {
			t.Errorf("got %v, want %v", gotNode.val, wantValue)
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

			if gotNode.val != wantVal {
				t.Errorf("got %v, want %v", gotNode.val, wantVal)
			}
		})
	}
}