package hash

import (
	"github.com/culionbear/qtool/nodes"
	"github.com/culionbear/qtool/qerror"
)

type cmdFunc func([][]byte) error

func (m cmdFunc) IsNil() bool {
	return m == nil
}

func (m *Manager) fetch(cmd uint8, info [][]byte) error {
	if f := m.cmdTable[cmd]; f != nil {
		return f(info)
	}
	return qerror.NewString("cmd is not found")
}

func (m *Manager) fetchSet(info [][]byte) error {
	if len(info) != 3 {
		return qerror.NewString("info length is error")
	}
	f, err := nodes.Get(info[1])
	if err != nil {
		return err
	}
	n, err := f(nil)
	if err != nil {
		return err
	}
	err = n.Deseriallize(info[2])
	if err != nil {
		return err
	}
	return m.Set(info[0], n)
}

func (m *Manager) fetchSetX(info [][]byte) error {
	if len(info) != 3 {
		return qerror.NewString("info length is error")
	}
	f, err := nodes.Get(info[1])
	if err != nil {
		return err
	}
	n, err := f(nil)
	if err != nil {
		return err
	}
	err = n.Deseriallize(info[2])
	if err != nil {
		return err
	}
	m.SetX(info[0], n)
	return nil
}

func (m *Manager) fetchUpdate(info [][]byte) error {
	if len(info) != 3 {
		return qerror.NewString("info length is error")
	}
	f, err := nodes.Get(info[1])
	if err != nil {
		return err
	}
	n, err := f(nil)
	if err != nil {
		return err
	}
	err = n.Deseriallize(info[2])
	if err != nil {
		return err
	}
	return m.Update(info[0], n)
}

func (m *Manager) fetchDel(info [][]byte) error {
	if len(info) != 1 {
		return qerror.NewString("info length is error")
	}
	return m.Del(info[0])
}

func (m *Manager) fetchDels(info [][]byte) error {
	if len(info) == 0 {
		return qerror.NewString("info length is error")
	}
	m.Dels(info...)
	return nil
}

func (m *Manager) fetchRename(info [][]byte) error {
	if len(info) != 2 {
		return qerror.NewString("info length is error")
	}
	return m.Rename(info[0], info[1])
}

func (m *Manager) fetchCover(info [][]byte) error {
	if len(info) != 2 {
		return qerror.NewString("info length is error")
	}
	return m.Cover(info[0], info[1])
}
