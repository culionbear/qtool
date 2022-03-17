package persistence

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

type aofModule struct {
	Cmd		uint8
	Values	[]interface{}
}

func newAofModule(cmd uint8, values []interface{}) *aofModule {
	return &aofModule{
		Cmd:	cmd,
		Values:	values,
	}
}

func (m *aofModule) getSetModule() ([]byte, error) {
	if len(m.Values) != 2 {
		return nil, qerror.NewString("values length is error")
	}
	k, ok := m.Values[0].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	n, ok := m.Values[1].(template.Node)
	if !ok {
		return nil, qerror.NewString("value type is not node")
	}
	return serializeNodeModule(CmdSet, k, n), nil
}

func (m *aofModule) getSetXModule() ([]byte, error) {
	if len(m.Values) != 2 {
		return nil, qerror.NewString("values length is error")
	}
	k, ok := m.Values[0].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	n, ok := m.Values[1].(template.Node)
	if !ok {
		return nil, qerror.NewString("value type is not node")
	}
	return serializeNodeModule(CmdSetX, k, n), nil
}

func (m *aofModule) getUpdateModule() ([]byte, error) {
	if len(m.Values) != 2 {
		return nil, qerror.NewString("values length is error")
	}
	k, ok := m.Values[0].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	n, ok := m.Values[1].(template.Node)
	if !ok {
		return nil, qerror.NewString("value type is not node")
	}
	return serializeNodeModule(CmdUpdate, k, n), nil
}


func (m *aofModule) getDelModule() ([]byte, error) {
	if len(m.Values) != 1 {
		return nil, qerror.NewString("values length is error")
	}
	k, ok := m.Values[0].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	return serializeBytesModule(CmdDel, k), nil
}

func (m *aofModule) getDelsModule() ([]byte, error) {
	list := make([][]byte, len(m.Values))
	for k, v := range m.Values {
		buf, ok := v.([]byte)
		if !ok {
			return nil, qerror.NewString("key type is not bytes")
		}
		list[k] = buf
	}
	return serializeBytesModule(CmdDel, list...), nil
}

func (m *aofModule) getRenameModule() ([]byte, error) {
	if len(m.Values) != 2 {
		return nil, qerror.NewString("values length is error")
	}
	k1, ok := m.Values[0].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	k2, ok := m.Values[1].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	return serializeBytesModule(CmdRename, k1, k2), nil
}

func (m *aofModule) getCoverModule() ([]byte, error) {
	if len(m.Values) != 2 {
		return nil, qerror.NewString("values length is error")
	}
	k1, ok := m.Values[0].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	k2, ok := m.Values[1].([]byte)
	if !ok {
		return nil, qerror.NewString("key type is not bytes")
	}
	return serializeBytesModule(CmdCover, k1, k2), nil
}