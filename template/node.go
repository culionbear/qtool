package template

import (
	"github.com/culionbear/qtool/ds/queue"
	"github.com/culionbear/qtool/qerror"
)

type Node interface {
	Object
	Type() []byte
	Serialize() []byte
	Deseriallize([]byte) *qerror.Error
}

type NewNode func(*queue.Manager[any]) (Node, *qerror.Error)

func (m NewNode) IsNil() bool {
	return m == nil
}
