package hash

import (
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
	for n := m.head.next; n != nil; n = n.next {
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
	for n := m.head.next; n != nil; n = n.next {
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
	for n := m.head.next; n != nil; n = n.next {
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

func (m *list) del(key []byte) (bool, qerror.Error) {
	if compare(m.head.key, key) {
		if m.head.next == nil {
			return true, nil
		}
		m.head = m.head.next
		m.head.last = nil
	} else if compare(m.tail.key, key) {
		m.tail = m.tail.last
		m.tail.next = nil
	} else {
		for n := m.head.next; n != m.tail; n = n.next {
			if compare(n.key, key) {
				n.last.next = n.next
				n.next.last = n.last
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
		m.tail = m.tail.last
		m.tail.next = nil
	} else {
		n.last.next = n.next
		n.next.last = n.last
	}
	return false
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
		if n.code&cap == 0 {
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
	var lo, hi *list
	if loHead != nil {
		loTail.next = nil
		lo = &list{
			head: loHead,
			tail: loTail,
		}
	}
	if hiHead != nil {
		hiTail.next = nil
		hi = &list{
			head: hiHead,
			tail: hiTail,
		}
	}
	return lo, hi
}
