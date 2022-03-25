package nodes

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

var m *Manager

func init() {
	m = New()
}

//Set node in trie
func Set(name []byte, f template.NewNode) qerror.Error {
	return m.Set(name, f)
}

//Get node in trie
func Get(buf []byte) (template.NewNode, qerror.Error) {
	return m.Get(buf)
}

func Exists(buf []byte) bool {
	return m.Exists(buf)
}

//List node name
func List() [][]byte {
	return m.list
}
