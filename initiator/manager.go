package initiator

import (
	"bytes"

	"github.com/culionbear/qtool/classes"
	"github.com/culionbear/qtool/ds/queue"
	"github.com/culionbear/qtool/protocol"
	"github.com/culionbear/qtool/qerror"
)

type manager struct {
	ch		chan *Module
	chClose	chan bool
	chWait	chan bool
	handler	*protocol.Manager
}

func newManager() *manager {
	return &manager{
		ch: make(chan *Module, 100000),
		chClose: make(chan bool),
		chWait: make(chan bool),
		handler: protocol.New(),
	}
}

func (m *manager) close() {
	m.chClose <- true
	<- m.chWait
}

func (m *manager) run() {
	defer func() {
		m.chWait <- true
	}()
	for {
		select {
		case ctx := <- m.ch:
			writer := &bytes.Buffer{}
			for ctx.list.Size() != 0 {
				writer.Write(
					m.execute(ctx.list.Pop()),
				)
			}
			ctx.Slot(m.handler.Write(writer.Bytes()))
		case <- m.chClose:
			return
		}
	}
}

func (m *manager) execute(msg *protocol.CmdTree) []byte {
	response := msg.Do(m.do)
	return m.handler.Pack(response)
}

func (m *manager) set(msg *Module) {
	m.ch <- msg
}

func (m *manager) do(cmd *queue.Manager[any]) any {
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
