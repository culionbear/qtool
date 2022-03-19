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
	VarList   []VarType
	ReturnVar []VarType
}

type Response struct {
	Err qerror.Error
	Msg []byte
}

type Request struct {
	Values []any
	Cmd    []byte
}

type Class interface {
	Object
	Do(*Request) *Response
	ApiList() []Api
}
