package protocol

type VarType byte

const (
	stringByte byte	= '='
	listByte byte	= '+'
	errorByte byte	= '!'
	numberByte byte	= '&'
	boolByte byte	= '?'
	floatByte byte	= '.'
)

const (
	separator uint8	= 0xff
	modValue		= 0xff
)

//Manager to protocol
type Manager struct {
	size	uint64
	buf		[]byte
}

//New Manager
func New() *Manager {
	return &Manager{}
}
