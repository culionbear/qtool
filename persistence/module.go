package persistence

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

type moduleFunc func(*module) ([]byte, *qerror.Error)

func (m moduleFunc) IsNil() bool {
	return m == nil
}

type module struct {
	Cmd    uint8
	Values []any
}

func newModule(cmd uint8, values []any) *module {
	return &module{
		Cmd:    cmd,
		Values: values,
	}
}

func (m *Manager) getSetModule(msg *module) ([]byte, *qerror.Error) {
	if len(msg.Values) != 2 {
		return nil, qerror.NewString("values length is error")
	}
	k, ok := msg.Values[0].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	n, ok := msg.Values[1].(template.Node)
	if !ok {
		return nil, qerror.NewString("value type is not node")
	}
	return serializeNodeModule(CmdSet, k, n), nil
}

func (m *Manager) getSetXModule(msg *module) ([]byte, *qerror.Error) {
	if len(msg.Values) != 2 {
		return nil, qerror.NewString("values length is error")
	}
	k, ok := msg.Values[0].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	n, ok := msg.Values[1].(template.Node)
	if !ok {
		return nil, qerror.NewString("value type is not node")
	}
	return serializeNodeModule(CmdSetX, k, n), nil
}

func (m *Manager) getUpdateModule(msg *module) ([]byte, *qerror.Error) {
	if len(msg.Values) != 2 {
		return nil, qerror.NewString("values length is error")
	}
	k, ok := msg.Values[0].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	n, ok := msg.Values[1].(template.Node)
	if !ok {
		return nil, qerror.NewString("value type is not node")
	}
	return serializeNodeModule(CmdUpdate, k, n), nil
}

func (m *Manager) getDelModule(msg *module) ([]byte, *qerror.Error) {
	if len(msg.Values) != 1 {
		return nil, qerror.NewString("values length is error")
	}
	k, ok := msg.Values[0].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	return serializeBytesModule(CmdDel, k), nil
}

func (m *Manager) getDelsModule(msg *module) ([]byte, *qerror.Error) {
	if len(msg.Values) != 0 {
		return nil, qerror.NewString("values length is error")
	}
	list := make([][]byte, len(msg.Values))
	for k, v := range msg.Values {
		buf, ok := v.([]byte)
		if !ok {
			return nil, qerror.NewString("key type is not bytes")
		}
		list[k] = buf
	}
	return serializeBytesModule(CmdDel, list...), nil
}

func (m *Manager) getRenameModule(msg *module) ([]byte, *qerror.Error) {
	if len(msg.Values) != 2 {
		return nil, qerror.NewString("values length is error")
	}
	k1, ok := msg.Values[0].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	k2, ok := msg.Values[1].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	return serializeBytesModule(CmdRename, k1, k2), nil
}

func (m *Manager) getCoverModule(msg *module) ([]byte, *qerror.Error) {
	if len(msg.Values) != 2 {
		return nil, qerror.NewString("values length is error")
	}
	k1, ok := msg.Values[0].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	k2, ok := msg.Values[1].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	return serializeBytesModule(CmdCover, k1, k2), nil
}
