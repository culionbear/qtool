package trie

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

type Manager[T template.Object] struct {
	head *node[T]
}

func New[T template.Object]() *Manager[T] {
	return &Manager[T]{
		head: newNode[T](),
	}
}

func (m *Manager[T]) Set(buf []byte, v T) qerror.Error {
	return m.head.add(buf, 0, len(buf), v)
}

func (m *Manager[T]) Get(buf []byte) (T, qerror.Error) {
	return m.head.get(buf, 0, len(buf))
}
