package qerror

import (
	"reflect"
	"unsafe"
)

// import (
// 	"reflect"
// 	"unsafe"
// )

//Error qlite system error type
type Error struct {
	msg []byte
}

//New error
func New(buf []byte) *Error {
	return &Error{
		buf,
	}
}

//Copy an new error
func (e *Error) Copy() *Error {
	buf := make([]byte, len(e.msg))
	copy(buf, e.msg)
	return &Error{
		buf,
	}
}

//New error to copy byte array
func Copy(src []byte) *Error {
	buf := make([]byte, len(src))
	copy(buf, src)
	return &Error{
		buf,
	}
}

//NewString new error to string
func NewString(msg string) *Error {
	b := make([]byte, len(msg))
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&msg))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return &Error{
		msg: b,
	}
}

//CopyError to Error
func CopyError(err error) *Error {
	if err == nil {
		return nil
	}
	return NewString(err.Error())
}

//Error to error interface
func (e *Error) Error() string {
	return string(e.msg)
}

func (e *Error) Size() int {
	return len(e.msg)
}

func (e *Error) Msg() []byte {
	return e.msg
}
