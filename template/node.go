package template

type Node interface {
	Object
	Type() []byte
	Serialize() []byte
	Deseriallize([]byte)
}

type NewNode func([]byte) (Node, error)

func (m NewNode) IsNil() bool {
	return m == nil
}
