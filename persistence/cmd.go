package persistence

import "github.com/culionbear/qtool/template"

const (
	CmdSet uint8 = 1 << iota
	CmdSetX
	CmdUpdate
	CmdDel
	CmdDels
	CmdRename
	CmdCover
)

//CmdOpt save options with cmd
type CmdOpt[T template.Object] [256]T

//NewCmd with empty array
func NewCmd[T template.Object]() *CmdOpt[T] {
	return &CmdOpt[T]{}
}

//Set options with cmd
func (m *CmdOpt[T]) Set(cmd uint8, opt T) {
	m[cmd] = opt
}

//Get options with cmd
func (m *CmdOpt[T]) Get(cmd uint8) T {
	return m[cmd]
}
