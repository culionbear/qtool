package queue

type Manager[T any] struct {
	head         *Node[T]
	tail         *Node[T]
	defaultValue *Node[T]
	size         int
}

func New[T any]() *Manager[T] {
	return &Manager[T]{
		defaultValue: &Node[T]{},
	}
}

func (m *Manager[T]) Push(v T) *Node[T] {
	n := newNode(v)
	if m.head == nil {
		m.head = n
		m.tail = n
	} else {
		m.tail.next = n
		m.tail = m.tail.next
	}
	m.size++
	return n
}

func (m *Manager[T]) Pop() T {
	if m.head == nil {
		return m.defaultValue.value
	}
	defer func() {
		m.head = m.head.next
		m.size--
	}()
	return m.head.value
}

func (m *Manager[T]) AutoIncrement() *Node[T] {
	n := &Node[T]{}
	if m.head == nil {
		m.head = n
		m.tail = n
	} else {
		m.tail.next = n
		m.tail = m.tail.next
	}
	m.size++
	return n
}

func (m *Manager[T]) Size() int {
	return m.size
}
