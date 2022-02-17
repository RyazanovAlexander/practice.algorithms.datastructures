package ordered

import (
	"constraints"
	"practice/data/structures/collection"
)

type comparer[T constraints.Ordered] struct{}

func (c comparer[T]) Less(x, y T) bool  { return x < y }
func (c comparer[T]) Equal(x, y T) bool { return x == y }

func Create[T constraints.Ordered]() collection.Comparer[T] {
	return comparer[T]{}
}
