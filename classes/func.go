package classes

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

var m *Manager

func init() {
	m = New()
}

//Set class in trie
func Set(name []byte, f template.Class) qerror.Error {
	return m.Set(name, f)
}

//Get class in trie
func Get(name []byte) (template.Class, qerror.Error) {
	return m.Get(name)
}

//List class name
func List() [][]byte {
	return m.list
}
