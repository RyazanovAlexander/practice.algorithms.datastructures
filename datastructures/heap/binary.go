package heap

import (
	"practice/primitives"
)

type binary[T any] struct {
	value T
}

func Create[T any](data []T, comparator primitives.Comparator[T]) (b *binary[T]) {
	return nil
}

func (b *binary[T]) Peek() (value T) {
	return b.value
}

func (b *binary[T]) Extract() (value T) {
	return b.value
}

func (b *binary[T]) Push() {

}
