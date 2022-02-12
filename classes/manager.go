package classes

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
	"github.com/culionbear/qtool/trie"
)

type Manager struct {
	tree *trie.Manager
	list [][]byte
}

func New() *Manager {
	return &Manager{
		tree: trie.New(),
		list: make([][]byte, 0),
	}
}

func (m *Manager) Set(str string, v template.Class) qerror.Error {
	err := m.tree.Set(str, v)
	if err != nil {
		return err
	}
	m.list = append(m.list, []byte(str))
	return nil
}

func (m *Manager) Get(buf []byte) (template.Class, qerror.Error) {
	v, err := m.tree.Get(buf)
	if err != nil {
		return nil, err
	}
	return v.(template.Class), nil
}

func (m *Manager) List() [][]byte {
	list := make([][]byte, len(m.list))
	for k, v := range m.list {
		buf := make([]byte, len(v))
		copy(buf, v)
		list[k] = buf
	}
	return list
}