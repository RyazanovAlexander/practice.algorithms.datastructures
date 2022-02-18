package comparer

import (
	"constraints"
	"practice/data/structures/collection"
	"practice/data/structures/collection/order"
)

type comparer[T constraints.Ordered] struct {
	order order.Order
}

func (c comparer[T]) Ordered(x, y T) bool { return c.order == order.Desc && x < y }
func (c comparer[T]) Equal(x, y T) bool   { return x == y }

func Descending[T constraints.Ordered]() collection.Orderer[T] {
	return comparer[T]{order: order.Desc}
}

func Ascending[T constraints.Ordered]() collection.Orderer[T] {
	return comparer[T]{order: order.Asc}
}
