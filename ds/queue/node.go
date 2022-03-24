package queue

type Node[T any] struct {
	value	T
	next	*Node[T]
}

func newNode[T any](value T) *Node[T] {
	return &Node[T]{
		value: value,
	}
}

func (n *Node[T]) SetValue(v T) {
	n.value = v
}
