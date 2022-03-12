package template

type Node interface {
	Type() []byte
	Serialize() []byte
	Deseriallize([]byte)
}

type NewNode func([]byte) Node
