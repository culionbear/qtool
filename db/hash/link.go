package hash

import "github.com/culionbear/qtool/template"

type link struct {
	size	int
	head	*node
}

func newLink(code uint32, key []byte, value template.Node) *link {
	return &link{
		size: 1,
		head: newNode(key, code, value),
	}
}

func (m *link) Set(code uint32, key []byte, value template.Node) {

}