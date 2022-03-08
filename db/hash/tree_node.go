package hash

import (
	"github.com/culionbear/qtool/template"
)

const(
	red		= true
	black	= false
)

type treeNode struct {
	key		[]byte
	value	template.Node
	code	uint32
	next	*treeNode

	color	bool
	parent	*treeNode
	left	*treeNode
	right	*treeNode
	prev	*treeNode
}

func newTreeode(key []byte, code uint32, value template.Node) *treeNode {
	return &treeNode{
		key: key,
		value: value,
		code: code,
	}
}

func (m *treeNode) Value() template.Node {
	return m.value
}

func (m *treeNode) SetValue(v template.Node) {
	m.value = v
}

func (m *treeNode) Key() []byte {
	k := make([]byte, len(m.key) )
	copy(k, m.key)
	return k
}

func (m *treeNode) Hex() uint32 {
	return m.code
}

func (m *treeNode) unsafeGetKey() []byte {
	return m.key
}

func (m *treeNode) Next() Node {
	return m.next
}