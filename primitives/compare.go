package primitives

import (
	value "practice/primitives/compared"
)

type Comparable[T any] interface {
	CompareTo(value T) value.Compared
}

type Comparator[T any] func(left, right T) value.Compared
