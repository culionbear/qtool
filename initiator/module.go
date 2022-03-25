package initiator

import (
	"github.com/culionbear/qtool/ds/queue"
	"github.com/culionbear/qtool/protocol"
)

//Module is a package to sand manager
type Module struct {
	callBack	chan []byte
	list		*queue.Manager[*protocol.CmdTree]
}

//NewModule with callback channel
func NewModule(ch chan[]byte) *Module {
	return &Module{
		callBack: ch,
	}
}

//CallBack return a byte array when option is ending
func (m *Module) CallBack() []byte {
	return <- m.callBack
}

//Set CmdTree in Module
func (m *Module) Set(cmd *queue.Manager[*protocol.CmdTree]) *Module {
	m.list = cmd
	return m
}

//Slot msg to other goroutine
func (m *Module) Slot(msg []byte) {
	m.callBack <- msg
}