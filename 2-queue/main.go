package main

import (
	"github.com/salmanrf/data-structures/queue/queue"
)

func main() {
	myq := queue.Circular[int](5)

	myq.Enqueue(100)
	myq.Enqueue(101)
	myq.Enqueue(102)
	myq.Enqueue(103)

	exists, item := myq.Dequeue()
	println("My front element", exists, item)

	exists, item = myq.Dequeue()
	println("My front element", exists, item)

	exists, item = myq.Dequeue()
	println("My front element", exists, item)

	exists, item = myq.Dequeue()
	println("My front element", exists, item)

	exists, item = myq.Dequeue()
	println("My front element", exists, item)
}