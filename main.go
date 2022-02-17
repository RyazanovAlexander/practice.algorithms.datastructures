package main

import (
	"fmt"

	"practice/data/structures/collection/orderer"
	"practice/data/structures/heap/binary"
)

func main() {
	data := []int{16, 11, 9, 10, 5, 6, 8, 1, 2, 4}
	heap := binary.Init(data, orderer.Create[int]())

	heap.Push(99)
	fmt.Print(heap.Pop())
	fmt.Print(heap.Peek())
	fmt.Print(heap.Len())

	// Add complex structure
	// Add 2 tests for demonstration
}
