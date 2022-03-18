package hash

import "github.com/culionbear/qtool/qerror"

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
	return nil
}

func (m *Manager) fetchSetX(info [][]byte) error {
	return nil
}

func (m *Manager) fetchUpdate(info [][]byte) error {
	return nil
}

func (m *Manager) fetchDel(info [][]byte) error {
	return nil
}

func (m *Manager) fetchDels(info [][]byte) error {
	return nil
}

func (m *Manager) fetchRename(info [][]byte) error {
	return nil
}

func (m *Manager) fetchCover(info [][]byte) error {
	return nil
}
