package protocol

type VarType byte

const (
	stringByte		= '='
	listByte		= '+'
	listValueByte	= '-'
	errorByte		= '!'
	numberByte		= '.'
	boolByte		= '?'
	cmdByte			= '#'
)

const (
	separator uint8	= 0xff
	modValue uint8	= 0xff
)

type Manager struct {
	size	uint64
	buf		[]byte
}

func New() *Manager {
	return &Manager{}
}