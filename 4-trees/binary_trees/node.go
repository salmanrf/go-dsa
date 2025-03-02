package binary_trees

type Node[T any] struct {
	Data T
	Left *Node[T]
	Right *Node[T]
}
