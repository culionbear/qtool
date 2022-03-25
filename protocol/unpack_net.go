package protocol

import (
	"github.com/culionbear/qtool/ds/queue"
	"github.com/culionbear/qtool/qerror"
)

func (m *Manager) UnpackNet(buf []byte) (*queue.Manager[*CmdTree], qerror.Error) {
	length := len(buf)
	list := queue.New[*CmdTree]()
	for i := 0; i < length; {
		sum, step, err := m.readSize(i, length, buf)
		if err != nil {
			return nil, err
		}
		i = step + sum
		if i > length {
			return nil, qerror.NewString("package length is error")
		}
		cmd, err := m.unpackNet(sum, buf[step:step+sum])
		if err != nil {
			return nil, err
		}
		list.Push(cmd)
	}
	return list, nil
}

func (m *Manager) unpackNet(length int, buf []byte) (*CmdTree, qerror.Error) {
	cmd := NewCmdTree(queue.New[any]())
	var err qerror.Error
	for i := 0; i < length; {
		if buf[i] == methodByte {
			sum, step, err := m.readSize(i+1, length, buf)
			if err != nil {
				return nil, err
			}
			i = step + sum
			if i > length {
				return nil, qerror.NewString("package length is error")
			}
			cCmd, err := m.unpackNet(sum, buf[step:step+sum])
			if err != nil {
				return nil, err
			}
			cmd.PushChild(cCmd)
		}
		i, err = m.unpack(cmd.cmd, i, length, buf)
		if err != nil {
			return nil, err
		}
	}
	return cmd, nil
}
