package trie

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

type node[T template.Object] struct {
	value    T
	flag	 bool
	children [26]*node[T]
}

func newNode[T template.Object]() *node[T] {
	return &node[T]{}
}

func (n *node[T]) add(buf []byte, i, length int, v T) *qerror.Error {
	if length == i {
		if !n.flag {
			return qerror.NewString("key is exists")
		}
		n.value = v
		n.flag = true
		return nil
	}
	k := int(buf[i] - 'a')
	if k < 0 || k >= 26 {
		return qerror.NewString("the key must be lowercase")
	}
	if n.children[k] == nil {
		n.children[k] = newNode[T]()
	}
	return n.children[k].add(buf, i+1, length, v)
}

func (n *node[T]) get(buf []byte, i, length int) (T, *qerror.Error) {
	if length == i {
		if n.flag {
			return n.value, qerror.NewString("key is not found")
		}
		return n.value, nil
	}
	k := int(buf[i] - 'a')
	if k < 0 || k >= 26 {
		return n.value, qerror.NewString("the key must be lowercase")
	}
	if n.children[k] == nil {
		return n.value, qerror.NewString("key is not found")
	}
	return n.children[k].get(buf, i+1, length)
}

func (n *node[T]) exists(buf []byte, i, length int) bool {
	if length == i {
		return !n.flag
	}
	k := int(buf[i] - 'a')
	if k < 0 || k >= 26 {
		return false
	}
	if n.children[k] == nil {
		return false
	}
	return n.children[k].exists(buf, i+1, length)
}
