package hash

import (
	"regexp"

	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

type list struct {
	head, tail *node
}

func newList(key []byte, code uint32, value template.Node) *list {
	n := newNode(key, code, value, nil)
	return &list{
		head: n,
		tail: n,
	}
}

func newListWithNode(n *node) *list {
	n.right, n.left = nil, nil
	return &list{
		head: n,
		tail: n,
	}
}

func (m *list) set(key []byte, code uint32, value template.Node) *qerror.Error {
	if m.head.code == code && compare(m.head.key, key) {
		return qerror.New(append(key, []byte(" is exists")...))
	}
	for n := m.head.right; n != nil; n = n.right {
		if n.code == code && compare(n.key, key) {
			return qerror.New(append(key, []byte(" is exists")...))
		}
	}
	m.tail.right = newNode(key, code, value, m.tail)
	m.tail = m.tail.right
	return nil
}

func (m *list) setX(key []byte, code uint32, value template.Node) int {
	if m.head.code == code && compare(m.head.key, key) {
		m.head.value = value
		return 0
	}
	for n := m.head.right; n != nil; n = n.right {
		if n.code == code && compare(n.key, key) {
			n.value = value
			return 0
		}
	}
	m.tail.right = newNode(key, code, value, m.tail)
	m.tail = m.tail.right
	return 1
}

func (m *list) update(key []byte, code uint32, value template.Node) *qerror.Error {
	if m.head.code == code && compare(m.head.key, key) {
		m.head.value = value
		return nil
	}
	for n := m.head.right; n != nil; n = n.right {
		if n.code == code && compare(n.key, key) {
			n.value = value
			return nil
		}
	}
	return qerror.New(append(key, []byte(" is not found")...))
}

func (m *list) get(key []byte) *node {
	if compare(m.head.key, key) {
		return m.head
	}
	for n := m.head.right; n != nil; n = n.right {
		if compare(n.key, key) {
			return n
		}
	}
	return nil
}

func (m *list) del(key []byte) (bool, *qerror.Error) {
	if compare(m.head.key, key) {
		if m.head.right == nil {
			return true, nil
		}
		m.head = m.head.right
		m.head.left = nil
	} else if compare(m.tail.key, key) {
		m.tail = m.tail.left
		m.tail.right = nil
	} else {
		for n := m.head.right; n != m.tail; n = n.right {
			if compare(n.key, key) {
				n.left.right = n.right
				n.right.left = n.left
				return false, nil
			}
		}
		return false, qerror.New(append(key, []byte(" is not found")...))
	}
	return false, nil
}

func (m *list) delNode(n *node) bool {
	if m.head == n {
		return true
	} else if m.tail == n {
		m.tail = m.tail.left
		m.tail.right = nil
	} else {
		n.left.right = n.right
		n.right.left = n.left
	}
	return false
}

func (m *list) pushBackNode(n *node) {
	n.left, n.right = m.tail, nil
	m.tail.right = n
	m.tail = n
}

func (m *list) iterators(f func(Node) bool) bool {
	if m.head == nil {
		return true
	}
	for n := m.head; n != nil; n = n.right {
		if !f(n) {
			return false
		}
	}
	return true
}

func (m *list) regexp(r *regexp.Regexp) [][]byte {
	if m.head == nil {
		return [][]byte{}
	}
	l := [][]byte{}
	for n := m.head; n != nil; n = n.right {
		l = append(l, n.key)
	}
	return l
}

func (m *list) onlyOne() bool {
	return m.head.right == nil
}

func (m *list) getHeadCode() uint32 {
	return m.head.code
}

func (m *list) resize(cap uint32) (store, store) {
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
