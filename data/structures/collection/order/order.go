package order

import (
	"constraints"
	"practice/data/structures/collection"
)

type Order int

const (
	Desc Order = iota
	Asc
)

type orderer[T constraints.Ordered] struct {
	order Order
}

func (c orderer[T]) Ordered(x, y T) bool { return c.order == Desc && x < y }

func Descending[T constraints.Ordered]() collection.Orderer[T] {
	return orderer[T]{order: Desc}
}

func Ascending[T constraints.Ordered]() collection.Orderer[T] {
	return orderer[T]{order: Asc}
}
