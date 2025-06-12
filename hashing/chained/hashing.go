package chained

import linkedlist "github.com/salmanrf/go-dsa/linked-list"

type hash_table struct {
	size int
	table []linkedlist.SingleLinkedList[int]
}

func compareFuncInt(a, b int) int {
	return b - a
}

type HashTable interface {
	Insert(key int)
	Get(key int) (bool, int)
}

func New(size int) HashTable {
	ht := &hash_table{
		size: size,
		table: make([]linkedlist.SingleLinkedList[int], size), 
	}

	return ht
}

func (ht *hash_table) hash_func(key int) int {
	return (key % ht.size)
}

func (ht *hash_table) Insert(key int) {
	index := ht.hash_func(key)

	if ht.table[index] == nil {
		ht.table[index] = linkedlist.New(key)
		return
	}

	chain := ht.table[index]

	chain.AppendSorted(key, compareFuncInt)
}

func (ht *hash_table) Get(key int) (bool, int) {
	index := ht.hash_func(key)

	chain := ht.table[index]

	if chain == nil {
		return false, 0
	}

	got := chain.Find(key, func (a, b int) bool { return a == b })

	if got == nil {
		return false, 0
	}

	return true, got.Val
}