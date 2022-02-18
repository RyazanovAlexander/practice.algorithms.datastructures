package collection

type Orderer[T any] interface {
	Ordered(x, y T) bool
}
