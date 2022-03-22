package protocol

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/culionbear/qtool/qerror"
)

func (m *Manager) Pack(module any) []byte {
	return m.pack(module)
}

func (m *Manager) pack(module any) []byte {
	switch msg := module.(type) {
	case string:
		return m.fromString(stringByte, m.s2b(msg))
	case []byte:
		return m.fromString(stringByte, msg)
	case qerror.Error:
		return m.fromString(errorByte, msg)
	case bool:
		return m.fromBool(msg)
	case int:
		return m.fromInt(msg)
	case []any:
		return m.fromList(msg)
	default:
		return m.fromString(errorByte, m.s2b("unknown type"))
	}
}

func (m *Manager) fromString(t byte, buf []byte) []byte {
	writer := &bytes.Buffer{}
	writer.WriteByte(t)
	size := len(buf)
	m.addNumber(writer, size)
	writer.WriteByte(separator)
	writer.Write(buf)
	return writer.Bytes()
}

func (m *Manager) fromBool(flag bool) []byte {
	if flag {
		return []byte{boolByte, 0x01, separator}
	}
	return []byte{boolByte, 0x00, separator}
}

func (m *Manager) fromInt(n int) []byte {
	writer := &bytes.Buffer{}
	writer.WriteByte(numberByte)
	m.addNumber(writer, n)
	writer.WriteByte(separator)
	return writer.Bytes()
}

func (m *Manager) fromFloat(float float64) []byte {
	writer := &bytes.Buffer{}
	writer.WriteByte(floatByte)
    bits := math.Float64bits(float)
    buf := []byte{0, 0, 0, 0, 0, 0, 0, 0, separator}
    binary.LittleEndian.PutUint64(buf, bits)
	writer.Write(buf)
    return writer.Bytes()
}

func (m *Manager) fromList(list []any) []byte {
	size := len(list)
	writer := &bytes.Buffer{}
	writer.WriteByte(listByte)
	m.addNumber(writer, size)
	writer.WriteByte(separator)
	for i := 0; i < size; i ++ {
		writer.Write(m.pack(list[i]))
	}
	return writer.Bytes()
}

func (m *Manager) addNumber(writer *bytes.Buffer, num int) {
	for num != 0 {
		writer.WriteByte(uint8(num % modValue))
		num /= modValue
	}
}