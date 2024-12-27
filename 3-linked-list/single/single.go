package single

type node[T any] struct {
	val T 
	next *node[T]
}

type linked_list[T any] struct {
	head *node[T]
}

type SingleLinkedList[T any] interface {
	Head() *node[T]
	Get(index int) (bool, *node[T])
	Insert(index int, val T) bool
	Append(T)
}

func New[T any](val T) (SingleLinkedList[T]) {
	head := &node[T] {val, nil}
	
	return &linked_list[T]{head}
}

func (l *linked_list[T]) Head() (*node[T]) {
	return l.head
}

func (l *linked_list[T]) Get(index int) (exists bool, nd *node[T]) {
	if index < 0 {
		return false, nd
	}
	
	i := 0
	node := l.head

	for node != nil {
		if i == index {
			return true, node
		}

		node = node.next
		i++
	}

	return false, nd
}

func (l *linked_list[T]) Append(val T) {
	curr := l.head
	i := 0
	
	for ; curr.next != nil; i++ {
		curr = curr.next
	}

	new := &node[T]{val, nil}

	curr.next = new
}

func (l *linked_list[T]) Insert(index int, val T) bool {
	if index < 0 {
		return false
	}

	new := node[T]{val: val, next: l.head}

	if index == 0 {
		l.head = &new
	}

	exists, nd := l.Get(index - 1)

	if !exists {
		return false
	}

	new.next = nd.next
	nd.next = &new

	return true
}