package orderer

import (
	"constraints"
	"practice/data/structures/collection"
)

type orderer[T constraints.Ordered] struct{}

func (c orderer[T]) Less(x, y T) bool { return x < y }

func Create[T constraints.Ordered]() collection.Orderer[T] {
	return orderer[T]{}
}
