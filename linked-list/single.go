package linkedlist

type ListNode[T any] struct {
	Val T 
	next *ListNode[T]
}

type linked_list[T any] struct {
	head *ListNode[T]
}

type SingleLinkedList[T any] interface {
	Head() *ListNode[T]
	Get(index int) (bool, *ListNode[T])
	Find(T, func (a, b T) bool) *ListNode[T]
	Insert(index int, val T) bool
	Append(T)
	AppendSorted(T, func (a, b T) int)
	Traverse(trFunc func (nd *ListNode[T]))
}

func New[T any](val T) (SingleLinkedList[T]) {
	head := &ListNode[T] {Val: val, next: nil}
	
	return &linked_list[T]{head}
}

func (l *linked_list[T]) Head() (*ListNode[T]) {
	return l.head
}

func (l *linked_list[T]) Get(index int) (exists bool, nd *ListNode[T]) {
	if index < 0 {
		return false, nd
	}
	
	i := 0
	node := l.head

	for i < index && node.next != nil {
		node = node.next

		i++
	}

	if i == index {
		return true, node
	}

	return false, nd
}

func (l *linked_list[T]) Find(val T, cf func (a, b T) bool) *ListNode[T] {
	return find(l.head, val, cf)
}

func find[T any](node *ListNode[T], val T, cf func (a, b T) bool) *ListNode[T] {
	if node == nil {
		return nil
	}
	
	if cf(val, node.Val) {
		return node
	}

	return find(node.next, val, cf)
}

func (l *linked_list[T]) Append(val T) {
	curr := l.head
	i := 0
	
	for ; curr.next != nil; i++ {
		curr = curr.next
	}

	new := &ListNode[T]{Val: val, next: nil}

	curr.next = new
}

func (l *linked_list[T]) AppendSorted(val T, compareFunc func (a, b T) int) {
	node := l.head

	new := &ListNode[T]{
		Val: val,
		next: nil,
	}
	
	if node == nil {
		l.head = new

		return
	}

	comp := compareFunc(node.Val, val)

	if comp < 0 {
		new.next = l.head
		l.head = new

		return
	}

	var prev *ListNode[T] = nil

	for node.next != nil {
		prev = node
		node = node.next
		comp = compareFunc(node.Val, val)

		if comp <= 0 {
			prev.next = new
			new.next = node
			
			return
		}
	}

	node.next = new
} 

func (l *linked_list[T]) Insert(index int, val T) bool {
	if index < 0 {
		return false
	}

	new := ListNode[T]{Val: val, next: l.head}

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

func (l *linked_list[T]) Traverse(trFunc func (nd *ListNode[T])) {
	n := l.head

	for n != nil {
		trFunc(n)

		n = n.next
	}
}