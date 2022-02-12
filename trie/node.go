package trie

import (
	"github.com/culionbear/qtool/qerror"
)

var(
	ErrKeyIsExists = qerror.Error("key is exists")
	ErrKeyIsNotFound = qerror.Error("key is not found")
	ErrKeyText = qerror.Error("the key must be lowercase")
)

type node struct {
	value interface{}
	children [26]*node
}

func newNode() *node{
	return &node{}
}

func (n *node) add(str string, i, length int, v interface{}) qerror.Error {
	if length == i {
		if n.value != nil {
			return ErrKeyIsExists
		}
		n.value = v
		return nil
	}
	k := int(str[i] - 'a')
	if k < 0 || k >= 26 {
		return ErrKeyText
	}
	if n.children[k] == nil {
		n.children[k] = newNode()
	}
	return n.children[k].add(str, i + 1, length, v)
}

func (n *node) get(buf []byte, i, length int) (interface{}, qerror.Error) {
	if length == i {
		if n.value == nil {
			return nil, ErrKeyIsNotFound
		}
		return n.value, nil
	}
	k := int(buf[i] - 'a')
	if k < 0 || k >= 26 {
		return nil, ErrKeyText
	}
	if n.children[k] == nil {
		n.children[k] = newNode()
	}
	return n.children[k].get(buf, i + 1, length)
}