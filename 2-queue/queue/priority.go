package queue

type priority_queue[T any] struct {
	size int
	front int
	rear int 
	container []T
}

type PriorityQueue[T any] interface {
	IsEmpty() bool
}

func Priority[T any](size int) PriorityQueue[T] {
	queue := &priority_queue[T]{
		size: size,
		front: -1,
		rear: -1,
		container: make([]T, size),
	}
	
	return queue
}

func (q *priority_queue[T]) IsEmpty() bool {
	return q.front == -1 && q.front == q.rear
}