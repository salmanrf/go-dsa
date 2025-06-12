package main

import "github.com/salmanrf/go-dsa/hashing"

func main() {
	hash_table := hashing.New(10)
	
	hash_table.Insert(777);
	hash_table.Insert(1092);
}