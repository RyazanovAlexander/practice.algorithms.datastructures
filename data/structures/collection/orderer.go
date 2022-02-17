package collection

type Orderer[T any] interface {
	Less(x, y T) bool
}
