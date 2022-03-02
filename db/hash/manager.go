package hash

import (
	"github.com/culionbear/qtool/db/iterator"
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

type Manager struct {

}

func New() *Manager {
	return &Manager{}
}

func (m *Manager) Set(key []byte, value template.Node) qerror.Error {
	return nil
}

func (m *Manager) SetX(key []byte, value template.Node) {
	
}

func (m *Manager) Get(key string) (template.Node, qerror.Error) {
	return nil, nil
}

func (m *Manager) Del(key... string) int {
	return 0
}

func (m *Manager) Regexp(str string) ([]template.Node, qerror.Error) {
	return nil, nil
}

func (m *Manager) Iterator(key []byte) (iterator.Node, qerror.Error) {
	return nil, nil
}

func (m *Manager) Rename(dst, src []byte) qerror.Error {
	return nil
}

func (m *Manager) Cover(dst, src []byte) {
	
}

func (m *Manager) Exist(key []byte) bool {
	return false
}

func (m *Manager) Size() int {
	return 0
}