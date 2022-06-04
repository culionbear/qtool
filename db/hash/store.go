package hash

import (
	"regexp"

	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

type store interface {
	set(key []byte, code uint32, value template.Node) *qerror.Error
	setX(key []byte, code uint32, value template.Node) int
	update(key []byte, code uint32, value template.Node) *qerror.Error
	get(key []byte) *node
	del(key []byte) (bool, *qerror.Error)
	delNode(n *node) bool
	pushBackNode(n *node)
	resize(cap uint32) (store, store)
	iterators(func(Node) bool) bool
	regexp(*regexp.Regexp) [][]byte
	onlyOne() bool
	getHeadCode() uint32
}