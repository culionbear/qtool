package initiator

import (
	"bytes"

	"github.com/culionbear/qtool/classes"
	"github.com/culionbear/qtool/ds/queue"
	"github.com/culionbear/qtool/protocol"
	"github.com/culionbear/qtool/qerror"
)

type Manager struct {
	ch      chan *Module
	chClose chan bool
	chWait  chan bool
	handler *protocol.Manager
}

func NewManager() *Manager {
	return &Manager{
		ch:      make(chan *Module, 100000),
		chClose: make(chan bool),
		chWait:  make(chan bool),
		handler: protocol.New(),
	}
}

func (m *Manager) Close() {
	m.chClose <- true
	<-m.chWait
}

func (m *Manager) Run() {
	defer func() {
		m.chWait <- true
	}()
	for {
		select {
		case ctx := <-m.ch:
			writer := &bytes.Buffer{}
			for ctx.list.Size() != 0 {
				writer.Write(
					m.Execute(ctx.list.Pop()),
				)
			}
			ctx.Slot(m.handler.Write(writer.Bytes()))
		case <-m.chClose:
			return
		}
	}
}

func (m *Manager) Execute(msg *protocol.CmdTree) []byte {
	response := msg.Do(m.Do)
	return m.handler.Pack(response)
}

func (m *Manager) Set(msg *Module) {
	m.ch <- msg
}

func (m *Manager) Do(cmd *queue.Manager[any]) any {
	if cmd.Size() < 2 {
		return qerror.NewString("cmd length is error")
	}
	className, ok := cmd.Pop().([]byte)
	if !ok {
		return qerror.NewString("class name is not string")
	}
	methodName, ok := cmd.Pop().([]byte)
	if !ok {
		return qerror.NewString("method name is not string")
	}
	classValue, err := classes.Get(className)
	if err != nil {
		return err
	}
	return classValue.Do(methodName, cmd)
}
