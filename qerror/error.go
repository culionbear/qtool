package qerror

type Error []byte

func (e Error) Copy() Error {
	buf := make(Error, len(e))
	copy(buf, e)
	return buf
}

func New(buf []byte) Error {
	return Error(buf)
}

func NewString(err string) Error {
	return Error(err)
}

func (e Error) Error() string {
	return string(e)
}