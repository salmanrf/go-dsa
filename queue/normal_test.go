package queue

import (
	"testing"
)

func TestNormalQueue(t *testing.T) {
  t.Run("Initiation", func (t *testing.T) {
    q := Normal[string](3)
    expectedExists, expectedValue := false, ""
  
    resultExists, resultValue := q.Peek()
  
    if resultExists != expectedExists || resultValue != expectedValue {
      t.Fatalf("Expected queue to be empty, got: %t, %s, want: %t, %s", resultExists, resultValue, expectedExists, expectedValue)
    }
  })

  t.Run("Enqueue", func (t *testing.T) {
    q := Normal[int](3)
    
    tests := []struct{
      name string
      input int
      want bool
    }{
      {"Item 1 of 3", 100, true},
      {"Item 2 of 3", 101, true},
      {"Item 3 of 3", 102, true},
      {"Item 4 of 3", 103, false},
      {"Item 4 of 3", 104, false},
    }

    for _, tt := range tests {
      t.Run(tt.name, func (t *testing.T) {
        result := q.Enqueue(tt.input)

        if result != tt.want {
          t.Fatalf("Expected return to be %t, got %t", tt.want, result)
        }
      })
    }
  })

  t.Run("Dequeue", func (t *testing.T) {
    q := Normal[int](3)
    values := []int{100, 101, 102}

    tests := []struct{
      name string
      wantExists bool
      wantValue int
    }{
      {"Item 1 of 3", true, values[0]},
      {"Item 2 of 3", true, values[1]},
      {"Item 3 of 3", true, values[2]},
      {"Item 4 of 3", false, 0},
    }

    q.Enqueue(values[0])
    q.Enqueue(values[1])
    q.Enqueue(values[2])

    for _, tt := range tests {
      t.Run(tt.name, func (t *testing.T) {
        resultExists, resultValue := q.Dequeue()

        if resultExists != tt.wantExists || resultValue != tt.wantValue {
          t.Fatalf("Expected dequeued element to be %t, %d, got %t, %d", tt.wantExists, tt.wantValue, resultExists, resultValue)
        }
      })
    }
  })
}