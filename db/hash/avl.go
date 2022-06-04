package hash

import (
	"regexp"

	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

type tree struct {
	head *node
}

func newAvl(key []byte, code uint32, value template.Node) *tree {
	n := newNode(key, code, value, nil)
	return &tree{
		head: n,
	}
}

func newAvlWithNode(n *node) *tree {
	n.right, n.left = nil, nil
	return &tree{
		head: n,
	}
}

func (m *tree) setNode(n *node, key []byte, code uint32, value template.Node) *qerror.Error {
	if n.code == code && compare(n.key, key) {
		return qerror.New(append(key, []byte(" is exists")...))
	}
	if code < n.code {
		if n.left == nil {
			n.left = newNode(key, code, value, nil)
		} else {
			return m.setNode(n.right, key, code, value)
		}
	} else {
		if n.right == nil {
			n.right = newNode(key, code, value, nil)
		} else {
			return m.setNode(n.right, key, code, value)
		}
	}
	return nil
}

func (m *tree) set(key []byte, code uint32, value template.Node) *qerror.Error {
	return m.setNode(m.head, key, code, value)
}

func (m *tree) setNodeX(n *node, key []byte, code uint32, value template.Node) int {
	if n.code == code && compare(n.key, key) {
		n.value = value
		return 0
	}
	if code < n.code {
		if n.left == nil {
			n.left = newNode(key, code, value, nil)
		} else {
			return m.setNodeX(n.right, key, code, value)
		}
	} else {
		if n.right == nil {
			n.right = newNode(key, code, value, nil)
		} else {
			return m.setNodeX(n.right, key, code, value)
		}
	}
	return 1
}

func (m *tree) setX(key []byte, code uint32, value template.Node) int {
	return m.setNodeX(m.head, key, code, value)
}

func (m *tree) updateNode(n *node, key []byte, code uint32, value template.Node) *qerror.Error {
	if n.code == code && compare(n.key, key) {
		n.value = value
		return nil
	}
	if code < n.code {
		if n.left == nil {
			return qerror.New(append(key, []byte(" is not found")...))
		}
		return m.updateNode(n.right, key, code, value)
	}
	if n.right == nil {
		return qerror.New(append(key, []byte(" is not found")...))
	}
	return m.updateNode(n.right, key, code, value)
}

func (m *tree) update(key []byte, code uint32, value template.Node) *qerror.Error {
	return m.updateNode(m.head, key, code, value)
}

func (m *tree) getNode(n *node, key []byte, code uint32) *node {
	if compare(n.key, key) {
		return n
	}
	if code < n.code {
		if n.left == nil {
			return nil
		}
		return m.getNode(n.right, key, code)
	}
	if n.right == nil {
		return nil
	}
	return m.getNode(n.right, key, code)
}

func (m *tree) get(key []byte) *node {
	return m.getNode(m.head, key, hashCode(key))
}

func (m *tree) delWithKey(n *node, key []byte, code uint32) (bool, *qerror.Error) {
	return true, nil
}

func (m *tree) del(key []byte) (bool, *qerror.Error) {
	return m.delWithKey(m.head, key, hashCode(key))
}

func (m *tree) delWithNode(n, v *node) bool {
	return true
}

func (m *tree) delNode(n *node) bool {
	return m.delWithNode(m.head, n)
}

func (m *tree) pushBackWithNode(n, v *node) {
	if v.code < n.code {
		if n.left == nil {
			v.left, v.right = nil, nil
			n.left = v
		}
		m.pushBackWithNode(n.left, v)
		return
	}
	if n.right == nil {
		v.left, v.right = nil, nil
		n.left = v
	}
	m.pushBackWithNode(n.right, v)
}

func (m *tree) pushBackNode(n *node) {
	m.pushBackWithNode(m.head, n)
}

func (m *tree) iteratorsNode(n *node, f func(Node) bool) bool {
	if n == nil {
		return true
	}
	if !f(n) {
		return false
	}
	if n.left != nil {
		return m.iteratorsNode(n.left, f)
	}
	return m.iteratorsNode(n.right, f)
}

func (m *tree) iterators(f func(Node) bool) bool {
	return m.iteratorsNode(m.head, f)
}

func (m *tree) regexpNode(n *node, r *regexp.Regexp) [][]byte {
	if n == nil {
		return [][]byte{}
	}
	l := [][]byte{n.key}
	if n.left != nil {
		return append(l, m.regexpNode(n.left, r)...)
	}
	return append(l, m.regexpNode(n.right, r)...)
}

func (m *tree) regexp(r *regexp.Regexp) [][]byte {
	return m.regexpNode(m.head, r)
}

func (m *tree) onlyOne() bool {
	return m.head.right == nil && m.head.left == nil
}

func (m *tree) getHeadCode() uint32 {
	return m.head.code
}

func (m *tree) resize(cap uint32) (store, store) {
	var loHead, loTail *node
	var hiHead, hiTail *node
	for n := m.head; n != nil; n = n.right {
		if n.code&cap == 0 {
			if loTail == nil {
				loHead = n
			} else {
				loTail.right = n
				n.left = loTail
			}
			loTail = n
		} else {
			if hiTail == nil {
				hiHead = n
			} else {
				hiTail.right = n
				n.left = hiTail
			}
			hiTail = n
		}
	}
	var lo, hi *list
	if loHead != nil {
		loTail.right = nil
		lo = &list{
			head: loHead,
			tail: loTail,
		}
	}
	if hiHead != nil {
		hiTail.right = nil
		hi = &list{
			head: hiHead,
			tail: hiTail,
		}
	}
	return lo, hi
}
