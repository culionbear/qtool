package protocol

import "github.com/culionbear/qtool/ds/queue"

type CmdFunc func(*queue.Manager[any]) any

type CmdTree struct {
	brother          *CmdTree
	child            *CmdTree
	childBrotherTail *CmdTree
	cmd              *queue.Manager[any]
	point            *queue.Node[any]
}

func NewCmdTree(cmd *queue.Manager[any]) *CmdTree {
	return &CmdTree{
		cmd: cmd,
	}
}

func (m *CmdTree) PushCmd(v any) {
	m.cmd.Push(v)
}

func (m *CmdTree) PushChild(child *CmdTree) {
	child.point = m.cmd.AutoIncrement()
	if m.child == nil {
		m.child = child
		m.childBrotherTail = m.child
	} else {
		m.childBrotherTail.brother = child
		m.childBrotherTail = child
	}
}

func (m *CmdTree) Do(f CmdFunc) any {
	for p := m.child; p != nil; p = p.brother {
		m.child.point.SetValue(m.child.Do(f))
	}
	return f(m.cmd)
}
