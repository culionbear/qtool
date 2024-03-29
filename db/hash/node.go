package hash

import (
	"github.com/culionbear/qtool/template"
)

type node struct {
	key         []byte
	value       template.Node
	code        uint32
	left, right *node
}

func newNode(key []byte, code uint32, value template.Node, last *node) *node {
	return &node{
		key:   key,
		value: value,
		code:  code,
		left:  last,
	}
}

func (m *node) Value() template.Node {
	return m.value
}

func (m *node) SetValue(v template.Node) {
	m.value = v
}

func (m *node) Key() []byte {
	k := make([]byte, len(m.key))
	copy(k, m.key)
	return k
}

func (m *node) Hex() uint32 {
	return m.code
}

func (m *node) rename(key []byte, code uint32) {
	m.key = key
	m.code = code
}
