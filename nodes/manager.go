package nodes

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
	"github.com/culionbear/qtool/trie"
)

//node manager
type Manager struct {
	tree *trie.Manager
	list [][]byte
}

//New Manager
func New() *Manager {
	return &Manager{
		tree: trie.New(),
		list: make([][]byte, 0),
	}
}

//Set node in trie
func (m *Manager) Set(name []byte, v template.NewNode) qerror.Error {
	err := m.tree.Set(name, v)
	if err != nil {
		return err
	}
	m.list = append(m.list, name)
	return nil
}

//Get node in trie
func (m *Manager) Get(name []byte) (template.NewNode, qerror.Error) {
	v, err := m.tree.Get(name)
	if err != nil {
		return nil, err
	}
	return v.(template.NewNode), nil
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
