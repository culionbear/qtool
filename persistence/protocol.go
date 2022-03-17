package persistence

import "github.com/culionbear/qtool/template"

const (
	separator	= ';'
)

func serializeNodeModule(cmd uint8, key []byte, node template.Node) []byte {
	msg := []byte{cmd}
	msg = append(msg, pack(key)...)
	msg = append(msg, pack(node.Type())...)
	msg = append(msg, pack(node.Serialize())...)
	return msg
}

func serializeBytesModule(cmd uint8, buf... []byte) []byte {
	msg := []byte{cmd}
	for _, v := range buf {
		msg = append(msg, pack(v)...)
	}
	return msg
}

func pack(buf []byte) []byte {
	size := len(buf)
	msg := make([]byte, 0)
	for size != 0 {
		msg = append(msg, byte(size % 256))
		size /= 256
	}
	msg = append(msg, separator)
	return append(msg, buf...)
}