package nodes

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

var m *Manager

func init() {
	m = New()
}

func Set(str string, f template.NewNode) qerror.Error {
	return m.Set(str, f)
}

func Get(buf []byte) (template.NewNode, qerror.Error) {
	return m.Get(buf)
}

func List() [][]byte {
	return m.list
}