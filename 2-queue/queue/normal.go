package queue

type normal_queue[T any] struct {
	front int
	rear int
	container []T
}

type NormalQueue[T any] interface {
	IsFull() bool
	IsEmpty() bool
	Enqueue(item T) bool
	Dequeue() (bool, T)
	Peek() (bool, T)
}

func Normal[T any](s int) NormalQueue[T] {
	queue := &normal_queue[T]{
		front: -1,
		rear: -1,
		container: make([]T, s),
	}

	return queue
}

func (q *normal_queue[T]) IsEmpty() bool {
	return q.front == q.rear && q.front == -1
}

func (q *normal_queue[T]) IsFull() bool {
	return q.rear == len(q.container) - 1
}

func (q *normal_queue[T]) Enqueue(item T) bool {
	if q.IsFull() {
		return false
	}
	
	if q.IsEmpty() {
		q.front += 1
	}
	
	q.rear += 1
	
	q.container[q.rear] = item

	return true
}

func (q *normal_queue[T]) Dequeue() (exists bool, item T) {
	if q.IsEmpty() {
		return false, item
	}

	item = q.container[q.front]
	
	if q.front == q.rear {
		q.front = -1
		q.rear = -1

		return true, item
	}
	
	q.front += 1
	
	return true, item
}

func (q *normal_queue[T]) Peek() (exists bool, item T) {
	if q.IsEmpty() {
		return false, item
	}
	
	item = q.container[q.front]
	
	return true, item
}

