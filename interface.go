package sort

type Interface[T any] interface {
	Less(i, j int) bool
	~[]T
}
