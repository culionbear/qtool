package protocol

type VarType byte

const (
	byteByte byte = iota
	stringByte
	errorByte
	numberByte
	trueByte
	falseByte
	floatByte
	methodByte
)

const separatorByte = 5

const (
	listByte byte = 1 << (iota + 8 - separatorByte)
)

const (
	separator uint8 = 0xff
	modValue        = 0xff
)

//Manager to protocol
type Manager struct {
	size int
	mod  int
}

//New Manager
func New() *Manager {
	return &Manager{
		mod: 1,
	}
}

func (m *Manager) init() {
	m.mod, m.size = 1, 0
}
