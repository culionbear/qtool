package qerror

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
func NewString(err string) Error {
	return Error(err)
}

//Error to error interface
func (e Error) Error() string {
	return string(e)
}
