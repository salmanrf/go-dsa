package chained

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	ht := New(10)

	t.Run("Should be able to get, return empty", func (t *testing.T) {
		gotExists, gotValue := ht.Get(99)
		wantExists := false 
		wantValue := 0

		if gotExists != wantExists {
			t.Errorf("got %v, want %v", gotExists, wantExists)
		}

		if gotValue != wantValue {
			t.Errorf("got %v, want %v", gotValue, wantValue)
		}
	})
}

func TestInsert(t *testing.T) {
	ht := New(13)

	ht.Insert(10)
	ht.Insert(11)
	ht.Insert(12)
	ht.Insert(22)
	ht.Insert(13)

	tests := []struct{
		val int
		exists bool
	}{
		{10, true},
		{11, true},
		{12, true},
		{22, true},
		{13, true},
	}

	for _, tt := range(tests) {
		t.Run(fmt.Sprintf("Element %v should exists: %v", tt.val, tt.exists), func (t *testing.T) {
			gotExists, gotValue := ht.Get(tt.val)
			wantExists, wantValue := tt.exists, tt.val

			if gotExists != wantExists {
				t.Errorf("got %v, want %v", gotExists, wantExists)
			}

			if gotValue != wantValue {
				t.Errorf("got %v, want %v", gotValue, wantValue)
			}
		})
	}
}
