package hash

type Node interface {
	Key() []byte
	Hex() uint32
}
