package classes

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
	"github.com/culionbear/qtool/trie"
)

//class manager
type Manager struct {
	tree *trie.Manager[template.Class]
	list [][]byte
}

//New Manager
func New() *Manager {
	return &Manager{
		tree: trie.New[template.Class](),
		list: make([][]byte, 0),
	}
}

//Set class in trie
func (m *Manager) Set(name []byte, v template.Class) qerror.Error {
	err := m.tree.Set(name, v)
	if err != nil {
		return err
	}
	m.list = append(m.list, []byte(name))
	return nil
}

//Get class in trie
func (m *Manager) Get(name []byte) (template.Class, qerror.Error) {
	return m.tree.Get(name)
}

//List class name
func (m *Manager) List() [][]byte {
	list := make([][]byte, len(m.list))
	for k, v := range m.list {
		buf := make([]byte, len(v))
		copy(buf, v)
		list[k] = buf
	}
	return list
}
