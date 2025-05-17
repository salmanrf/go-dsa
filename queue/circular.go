package queue

type circular_queue[T any] struct {
	size int
	front int
	rear int
	container []T
}

type CircularQueue[T any] interface {
	IsFull() bool
	IsEmpty() bool
	Enqueue(item T) bool
	Dequeue() (bool, T)
	Peek() (bool, T)
}

func Circular[T any](s int) NormalQueue[T] {
	queue := &circular_queue[T]{
		front: -1,
		rear: -1,
		size: s,
		container: make([]T, s),
	}

	return queue
}

func (q *circular_queue[T]) IsEmpty() bool {
	return q.front == q.rear && q.front == -1
}

func (q *circular_queue[T]) IsFull() bool {
	return (q.front == 0 && q.rear == q.size - 1) || (q.front == q.rear + 1)
}

func (q *circular_queue[T]) Enqueue(item T) bool {
	if q.IsFull() {
		return false
	}
	
	if q.IsEmpty() {
		q.front = 0
		q.rear = 0
	} else {
		q.rear = (q.rear + 1) % q.size
	}
	
	q.container[q.rear] = item

	return true
}

func (q *circular_queue[T]) Dequeue() (exists bool, item T) {
	if q.IsEmpty() {
		return false, item
	}

	item = q.container[q.front]
	
	if q.front == q.rear {
		q.front = -1
		q.rear = -1

		return true, item
	}
	
	q.front = (q.front + 1) % q.size
	
	return true, item
}

func (q *circular_queue[T]) Peek() (exists bool, item T) {
	if q.IsEmpty() {
		return false, item
	}
	
	item = q.container[q.front]
	
	return true, item
}

