package trie

import "github.com/culionbear/qtool/qerror"

type Manager struct {
	head *node
}

func New() *Manager {
	return &Manager{
		head: newNode(),
	}
}

func (m *Manager) Set(buf []byte, v interface{}) qerror.Error {
	return m.head.add(buf, 0, len(buf), v)
}

func (m *Manager) Get(buf []byte) (interface{}, qerror.Error) {
	return m.head.get(buf, 0, len(buf))
}
