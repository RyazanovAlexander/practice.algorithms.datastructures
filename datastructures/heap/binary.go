package heap

type Interface[T any] interface {
	Push(x T)
	Pop() T
	Peek() T
	Len() int
}

type container[T any] struct {
	heap []T
}

func Init[T any](data []T) Interface[T] {
	return container[T]{
		heap: data,
	}
}

func (h container[T]) Push(x T) {

}

func (h container[T]) Pop() T {
	return h.heap[0]
}

func (h container[T]) Peek() T {
	return h.heap[0]
}

func (h container[T]) Len() int {
	return len(h.heap)
}
