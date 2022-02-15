package main

import (
	"fmt"

	"practice/datastructures/heap"
	"practice/primitives"
	value "practice/primitives/compared"
)

var comparator primitives.Comparator[int] = func(left, right int) value.Compared {
	if left < right {
		return value.Less
	}
	if left > right {
		return value.Larger
	}
	return value.Equal
}

func main() {
	data := []int{16, 11, 9, 10, 5, 6, 8, 1, 2, 4}
	binaryHeap := heap.Create(data, comparator)

	fmt.Println(binaryHeap.Peek())
}
