package protocol

type VarType byte

const (
	stringByte byte = iota
	errorByte
	numberByte
	trueByte
	falseByte
	floatByte
	listByte
	funcByte
)

const (
	chByte byte	= 1 << (iota + 3)
)

const (
	separator uint8 = 0xff
	modValue        = 0xff
	cmdValue		= 5
)

//Manager to protocol
type Manager struct {
	size	int
	mod		int
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
