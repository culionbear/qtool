package template

import "github.com/culionbear/qtool/qerror"

const (
	Number = iota
	String
	List
	Bool
)

type VarType int

type Api struct {
	VarList	[]VarType
	ReturnVar VarType
	ReadOnly bool
}

type Result struct {
	Err qerror.Error
	Msg []byte
	Persistence bool
}

type Class interface {
	Do([]byte) Result
	ApiList() []Api
}