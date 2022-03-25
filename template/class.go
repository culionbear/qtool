package template

import (
	"github.com/culionbear/qtool/ds/queue"
	"github.com/culionbear/qtool/protocol"
)

//Api is Class's api list node
//Need define with user
type Api struct {
	Name      []byte
	VarList   []protocol.VarType
	ReturnVar []protocol.VarType
}

type Class interface {
	Object
	Do([]byte, *queue.Manager[any]) any
	ApiList() []Api
	Name() []byte
}
