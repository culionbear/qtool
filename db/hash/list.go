package hash

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

type list struct {
	head, tail	*node
}

func newList(key []byte, code uint32, value template.Node) *list {
	n := newNode(key, code, value, nil)
	return &list{
		head: n,
		tail: n,
	}
}

func newListWithNode(n *node) *list {
	n.next, n.last = nil, nil
	return &list{
		head: n,
		tail: n,
	}
}

func (m *list) set(key []byte, code uint32, value template.Node) qerror.Error {
	if m.head.code == code && compare(m.head.key, key) {
		return qerror.New(append(key, []byte(" is exists")...))
	}
	for n := m.head.next; n != nil ; n = n.next {
		if n.code == code && compare(n.key, key) {
			return qerror.New(append(key, []byte(" is exists")...))
		}
	}
	m.tail.next = newNode(key, code, value, m.tail)
	m.tail = m.tail.next
	return nil
}

func (m *list) setX(key []byte, code uint32, value template.Node) int {
	if m.head.code == code && compare(m.head.key, key) {
		m.head.value = value
		return 0
	}
	for n := m.head.next; n != nil ; n = n.next {
		if n.code == code && compare(n.key, key) {
			n.value = value
			return 0
		}
	}
	m.tail.next = newNode(key, code, value, m.tail)
	m.tail = m.tail.next
	return 1
}

func (m *list) update(key []byte, code uint32, value template.Node) qerror.Error {
	if m.head.code == code && compare(m.head.key, key) {
		m.head.value = value
		return nil
	}
	for n := m.head.next; n != nil ; n = n.next {
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
	for n := m.head.next; n != nil; n = n.next {
		if compare(n.key, key) {
			return n
		}
	}
	return nil
}

func (m *list) pushBackNode(n *node) {
	n.last, n.next = m.tail, nil
	m.tail.next = n
	m.tail = n
}

func (m *list) resize(cap uint32) (*list, *list) {
	//TODO:跳表实现
	var loHead, loTail *node
	var hiHead, hiTail *node
	for n := m.head; n != nil; n = n.next {
		if n.code & cap == 0 {
			if loTail == nil {
				loHead = n
			} else {
				loTail.next = n
				n.last = loTail
			}
			loTail = n
		} else {
			if hiTail == nil {
				hiHead = n
			} else {
				hiTail.next = n
				n.last = hiTail
			}
			hiTail = n
		}
	}
	return &list{
		head: loHead,
		tail: loTail,
	}, &list{
		head: hiHead,
		tail: hiTail,
	}
}