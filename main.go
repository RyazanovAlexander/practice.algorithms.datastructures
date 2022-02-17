package main

import (
	"fmt"
	"practice/datastructures/heap"
)

func main() {
	data := []int{16, 11, 9, 10, 5, 6, 8, 1, 2, 4}
	binaryHeap := heap.Init(data)

	binaryHeap.Push(99)

	value := binaryHeap.Pop()
	fmt.Print(value)

	value = binaryHeap.Peek()
	fmt.Print(value)

	value = binaryHeap.Len()
	fmt.Print(value)
}
