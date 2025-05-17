package stack

import "fmt"

type stack[T any] struct {
	size int
	top int
	container []T
}

type Stack[T any] interface {
	Push(item T) 
	Pop() (bool, T)
	IsEmpty() bool
	IsFull() bool
	Peek() (bool, T)
}

func New[T any](size int) Stack[T] {
	stack := &stack[T]{container: make([]T, size), top: -1, size: size}

	return stack
}

func (s *stack[T]) Push(item T) {
	if s.IsFull() {
		return 
	}

	fmt.Println("Stack container size", len(s.container))
	fmt.Println("Top", s.top)

	s.top += 1
	s.container[s.top] = item
}

func (s *stack[T]) Pop() (exists bool, item T) {
	if s.IsEmpty() {
		return false, item
	}

	item = s.container[s.top]
	s.top -= 1

	return true, item
}

func (s *stack[T]) IsEmpty() bool {
	return s.top <= -1
}

func (s *stack[T]) IsFull() bool {
	return s.top + 1 >= s.size
}

func (s *stack[T]) Peek() (exists bool, item T) {
	if s.IsEmpty() {
		return false, item
	}
	
	return true, s.container[s.top]
}

