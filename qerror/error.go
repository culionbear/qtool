package qerror

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
	// /* #nosec G103 */
	// errH := (*reflect.SliceHeader)(unsafe.Pointer(&err))
	// /* #nosec G103 */
	// sh := (*reflect.StringHeader)(unsafe.Pointer(&msg))
	// errH.Data = sh.Data
	// errH.Cap = sh.Len
	// errH.Len = sh.Len
	return &Error{
		msg: []byte(msg),
	}
}

//CopyError to Error
func CopyError(err error) *Error {
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
