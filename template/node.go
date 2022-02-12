package template

type Node interface {
	Type() []byte
}

type NewNode func() Node