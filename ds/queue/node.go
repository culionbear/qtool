package queue

type node[T any] struct {
	value	T
	next	*node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		value: value,
	}
}
