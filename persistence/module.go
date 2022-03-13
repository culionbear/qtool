package persistence

import (
	"github.com/culionbear/qtool/db/hash"
)

//Node struct in binary file
type Node struct {
	Key   []byte
	Value hash.Node
}

//NewNode return Node point with k/v
func NewNode(key []byte, value hash.Node) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}
