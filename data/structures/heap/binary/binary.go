package binary

import (
	"practice/data/structures/collection"
)

type Interface[T any] interface {
	Push(x T)
	Pop() T
	Peek() T
	Len() int
}

type container[T any] struct {
	heap    []T
	orderer collection.Orderer[T]
}

func Heapify[T any](data []T, orderer collection.Orderer[T]) Interface[T] {
	return container[T]{
		heap:    data,
		orderer: orderer,
	}
}

func (c container[T]) Push(x T) {

}

func (c container[T]) Pop() T {
	return c.heap[0]
}

func (c container[T]) Peek() T {
	return c.heap[0]
}

func (c container[T]) Len() int {
	return len(c.heap)
}
