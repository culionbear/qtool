package nodes

import (
	"github.com/culionbear/qtool/template"
)

var m *Manager

func init() {
	m = New()
}

//Set node in trie
func Set(name []byte, f template.NewNode) error {
	return m.Set(name, f)
}

//Get node in trie
func Get(buf []byte) (template.NewNode, error) {
	return m.Get(buf)
}

//List node name
func List() [][]byte {
	return m.list
}
