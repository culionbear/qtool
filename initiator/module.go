package initiator

import "github.com/culionbear/qtool/protocol"

//Module is a package to sand manager
type Module struct {
	callBack	chan []byte
	list		[]protocol.CmdTree
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
func (m *Module) Set(cmd []protocol.CmdTree) {
	m.list = cmd
}

//Slot msg to other goroutine
func (m *Module) Slot(msg []byte) {
	m.callBack <- msg
}