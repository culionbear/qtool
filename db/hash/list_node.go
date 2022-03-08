package hash

import (
	"github.com/culionbear/qtool/template"
)

type listNode struct {
	key		[]byte
	value	template.Node
	code	uint32
	next	*listNode
}

func newListNode(key []byte, code uint32, value template.Node) *listNode {
	return &listNode{
		key: key,
		value: value,
		code: code,
	}
}

func (m *listNode) Value() template.Node {
	return m.value
}

func (m *listNode) SetValue(v template.Node) {
	m.value = v
}

func (m *listNode) Key() []byte {
	k := make([]byte, len(m.key) )
	copy(k, m.key)
	return k
}

func (m *listNode) Hex() uint32 {
	return m.code
}

func (m *listNode) unsafeGetKey() []byte {
	return m.key
}

func (m *listNode) Next() Node {
	return m.next
}