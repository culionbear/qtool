package hash

import "github.com/culionbear/qtool/template"

type Node interface {
	Value() template.Node
	SetValue(template.Node)
	Key() []byte
	Hex() uint32
}