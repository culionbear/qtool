package hash

import "github.com/culionbear/qtool/template"

type Node interface {
	Value() template.Node
	SetValue(template.Node)
	Key() []byte
	Hex() uint32
	Next() Node
	unsafeGetKey() []byte
}