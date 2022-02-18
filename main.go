package main

import (
	"fmt"

	"practice/data/structures/collection/order"
	"practice/data/structures/heap/binary"
)

func main() {
	tree := []int{16, 11, 9, 10, 5, 6, 8, 1, 2, 4}
	heap := binary.Heapify(tree, order.Descending[int]())

	heap.Push(99)
	fmt.Println(heap.Pop())
	fmt.Println(heap.Peek())
	fmt.Println(heap.Len())

	// Add complex structure
	// Add 2 tests for demonstration
}
