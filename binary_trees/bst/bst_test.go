package bst

import "testing"

func TestNew(t *testing.T) {
	data := 10

	want := data

	got := New(data, func(a, b int) int {
		return a - b
	}).Data

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestInsert(t *testing.T) {
	root := New(10, func(a, b int) int {
		return a - b
	})

	root.Insert(5)
	root.Insert(15)
	root.Insert(6)
	root.Insert(12)

	want := 5

	got := root.Left.Data

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	} 

	want = 6

	got = root.Left.Right.Data

	if got != want {
		t.Errorf("Got %d, want %d", got, want)
	}
}

func TestMin(t *testing.T) {
	root := New(10, func(a, b int) int {
		return a - b
	})

	root.Insert(5)
	root.Insert(15)
	root.Insert(6)
	root.Insert(12)

	want := 5

	got := root.Min()

	if got.Data != want {
		t.Errorf("Got %d, want %d", got.Data, want)
	}
}

func TestMax(t *testing.T) {
	tests := []struct{
		name string
		elements []int
		want int
	} {
		{
			"Case 1",
			[]int{10, 5, 15, 6, 12},
			15,
		},
		{
			"Case 2",
			[]int{100, 50, 25, 12, 6, 3},
			100,
		},
		{
			"Case 3",
			[]int{0, 1, 2, 3, 4, 5},
			5,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := New(tt.elements[0], func(a, b int) int {
				return a - b
			})

			for _, el := range tt.elements[1:] {
				root.Insert(el)
			}
			
			want := tt.want
		
			got := root.Max()
		
			if got.Data != want {
				t.Errorf("Got %d, want %d", got.Data, want)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	tests := []struct{
		name 				string
		elements  	[]int
		del 				int
		replacement int
	}{
		{
			"Delete Node with right subtree",
			[]int{1, 2, 3, 4, 5},
			3,
			4,
		},
		{
			"Delete Node with left & right subtree",
			[]int{20, 10, 30},
			20,
			30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			root := New(tt.elements[0], func(a, b int) int {
				return a - b
			})

			for _, el := range tt.elements[1:] {
				root.Insert(el)
			}
			
			root.Delete(tt.del)

			t.Run("The deleted node should not exist anymore", func (t *testing.T) {
				var want *BSTNode[int] = nil 
				got := root.Search(tt.del)
			
				if got != want {
					t.Errorf("Got %p, want %p", got, want)
				}
			})
		})
	}
}