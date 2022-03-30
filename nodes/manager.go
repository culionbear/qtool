package nodes

import (
	"github.com/culionbear/qtool/ds/trie"
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

//node manager
type Manager struct {
	tree *trie.Manager[template.NewNode]
	list [][]byte
}

//New Manager
func New() *Manager {
	return &Manager{
		tree: trie.New[template.NewNode](),
		list: make([][]byte, 0),
	}
}

//Set node in trie
func (m *Manager) Set(name []byte, v template.NewNode) *qerror.Error {
	err := m.tree.Set(name, v)
	if err != nil {
		return err
	}
	m.list = append(m.list, name)
	return nil
}

//Get node in trie
func (m *Manager) Get(name []byte) (template.NewNode, *qerror.Error) {
	return m.tree.Get(name)
}

func (m *Manager) Exists(name []byte) bool {
	return m.tree.Exists(name)
}

//List node name
func (m *Manager) List() [][]byte {
	list := make([][]byte, len(m.list))
	for k, v := range m.list {
		buf := make([]byte, len(v))
		copy(buf, v)
		list[k] = buf
	}
	return list
}
