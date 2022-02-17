package collection

type Comparer[T any] interface {
	Orderer[T]

	Equal(x, y T) bool
}
