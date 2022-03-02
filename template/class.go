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

type Response struct {
	Err qerror.Error
	Msg []byte
	Persistence bool
}

type Request struct {
	Value []interface{}
	Api []byte
}

type Class interface {
	Do(Request) Response
	ApiList() []Api
}