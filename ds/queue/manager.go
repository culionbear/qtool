package queue

type Manager[T any] struct {
	head			*node[T]
	tail			*node[T]
	defaultValue	*node[T]
	size	int
}

func New[T any]() *Manager[T] {
	return &Manager[T]{
		defaultValue: &node[T]{},
	}
}

func (m *Manager[T]) Push(v T) {
	if m.head == nil {
		m.head = newNode(v)
		m.tail = newNode(v)
	} else {
		m.tail.next = newNode(v)
		m.tail = m.tail.next
	}
	m.size ++
}

func (m *Manager[T]) Pop() T {
	if m.head == nil {
		return m.defaultValue.value
	}
	defer func() {
		m.head = m.head.next
		m.size --
	}()
	return m.head.value
}

func (m *Manager[T]) Size() int {
	return m.size
}
