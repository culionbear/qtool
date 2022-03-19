package persistence

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

const (
	headLength = 1 << 3
)

func serializeNodeModule(cmd uint8, key []byte, node template.Node) []byte {
	msg := []byte{cmd}
	msg = append(msg, pack(key)...)
	msg = append(msg, pack(node.Type())...)
	msg = append(msg, pack(node.Serialize())...)
	return pack(msg)
}

func serializeBytesModule(cmd uint8, buf ...[]byte) []byte {
	msg := []byte{cmd}
	for _, v := range buf {
		msg = append(msg, pack(v)...)
	}
	return pack(msg)
}

func pack(buf []byte) []byte {
	size := len(buf)
	msg := newBytes()
	for i := 0; size != 0 && i < headLength; i++ {
		msg[i] = byte(size % 256)
		size /= 256
	}
	return append(msg, buf...)
}

func getPackageLength(buf []byte) (uint64, error) {
	length := len(buf)
	if length < headLength {
		return 0, qerror.NewString("bytes length is error")
	}
	var size, num uint64 = 0, 1
	for i := 0; i < headLength; i++ {
		size += uint64(buf[i]) * num
		num *= 256
	}
	return size, nil
}

func newBytes() []byte {
	return []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
}
