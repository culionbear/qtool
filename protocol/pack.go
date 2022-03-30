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
	if module == nil {
		return m.fromBool(true)
	}
	switch msg := module.(type) {
	case byte:
		return m.fromByte(msg)
	case string:
		return m.fromString(m.s2b(msg))
	case *qerror.Error:
		return m.fromError(msg)
	case bool:
		return m.fromBool(msg)
	case int:
		return m.fromInt(msg)
	case float64:
		return m.fromFloat(msg)
	case []byte:
		return m.fromString(msg)
	case []string:
		return m.fromStringList(msg)
	case []*qerror.Error:
		return m.fromErrorList(msg)
	case []bool:
		return m.fromBoolList(msg)
	case []int:
		return m.fromIntList(msg)
	case []float64:
		return m.fromFloatList(msg)
	default:
		return m.fromError(qerror.NewString("unknown type"))
	}
}

func (m *Manager) fromByte(buf byte) []byte {
	return []byte{byteByte, buf}
}

func (m *Manager) fromString(buf []byte) []byte {
	writer := &bytes.Buffer{}
	writer.WriteByte(stringByte)
	size := len(buf)
	m.addNumber(writer, size)
	writer.Write(buf)
	return writer.Bytes()
}

func (m *Manager) fromError(buf *qerror.Error) []byte {
	writer := &bytes.Buffer{}
	writer.WriteByte(errorByte)
	size := buf.Size()
	m.addNumber(writer, size)
	writer.Write(buf.Msg())
	return writer.Bytes()
}

func (m *Manager) fromBool(flag bool) []byte {
	if flag {
		return []byte{trueByte}
	}
	return []byte{falseByte}
}

func (m *Manager) fromInt(n int) []byte {
	writer := &bytes.Buffer{}
	writer.WriteByte(numberByte)
	m.addNumber(writer, n)
	return writer.Bytes()
}

func (m *Manager) fromFloat(float float64) []byte {
	writer := &bytes.Buffer{}
	writer.WriteByte(floatByte)
	bits := math.Float64bits(float)
	buf := make([]byte, 8, 8)
	binary.LittleEndian.PutUint64(buf, bits)
	writer.Write(buf)
	return writer.Bytes()
}

func (m *Manager) addNumber(writer *bytes.Buffer, num int) {
	for num != 0 {
		writer.WriteByte(uint8(num % modValue))
		num /= modValue
	}
	writer.WriteByte(separator)
}

func (m *Manager) fromStringList(list []string) []byte {
	writer := &bytes.Buffer{}
	writer.WriteByte(listByte | stringByte)
	size := len(list)
	m.addNumber(writer, size)
	for i := 0; i < size; i++ {
		m.addNumber(writer, len(list[i]))
		writer.Write(m.s2b(list[i]))
	}
	return writer.Bytes()
}

func (m *Manager) fromErrorList(list []*qerror.Error) []byte {
	writer := &bytes.Buffer{}
	writer.WriteByte(listByte | errorByte)
	size := len(list)
	m.addNumber(writer, size)
	for i := 0; i < size; i++ {
		m.addNumber(writer, list[i].Size())
		writer.Write(list[i].Msg())
	}
	return writer.Bytes()
}

func (m *Manager) fromBoolList(list []bool) []byte {
	writer := &bytes.Buffer{}
	writer.WriteByte(listByte | trueByte)
	size := len(list)
	m.addNumber(writer, size)
	for i := 0; i < size; i++ {
		if list[i] {
			writer.WriteByte(trueByte)
		} else {
			writer.WriteByte(falseByte)
		}
	}
	return writer.Bytes()
}

func (m *Manager) fromIntList(list []int) []byte {
	writer := &bytes.Buffer{}
	writer.WriteByte(listByte | numberByte)
	size := len(list)
	m.addNumber(writer, size)
	for i := 0; i < size; i++ {
		m.addNumber(writer, list[i])
	}
	return writer.Bytes()
}

func (m *Manager) fromFloatList(list []float64) []byte {
	writer := &bytes.Buffer{}
	writer.WriteByte(listByte | floatByte)
	size := len(list)
	m.addNumber(writer, size)
	for i := 0; i < size; i++ {
		bits := math.Float64bits(list[i])
		buf := make([]byte, 8, 8)
		binary.LittleEndian.PutUint64(buf, bits)
		writer.Write(buf)
	}
	return writer.Bytes()
}
