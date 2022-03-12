package hash

import (
	"github.com/culionbear/qtool/template"
)

type node struct {
	key        []byte
	value      template.Node
	code       uint32
	last, next *node
}

func newNode(key []byte, code uint32, value template.Node, last *node) *node {
	return &node{
		key:   key,
		value: value,
		code:  code,
		last:  last,
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

func (m *node) resize(key []byte, code uint32) {
	m.deled()
	m.rename(key, code)
}

func (m *node) deled() {
	if m.last != nil {
		m.last.next = m.next
	}
	if m.next != nil {
		m.next.last = m.last
	}
}

func (m *node) rename(key []byte, code uint32) {
	m.key = key
	m.code = code
}
