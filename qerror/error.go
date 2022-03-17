package qerror

import (
	"reflect"
	"unsafe"
)

//Error qlite system error type
type Error []byte

//Copy an new error
func (e Error) Copy() Error {
	buf := make(Error, len(e))
	copy(buf, e)
	return buf
}

//New error
func New(buf []byte) Error {
	return Error(buf)
}

//NewString new error to string
func NewString(msg string) (err Error) {
	/* #nosec G103 */
	errH := (*reflect.SliceHeader)(unsafe.Pointer(&err))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&msg))
	errH.Data = sh.Data
	errH.Cap = sh.Len
	errH.Len = sh.Len
	return
}

//Error to error interface
func (e Error) Error() string {
	return string(e)
}
