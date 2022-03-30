package protocol

import "bytes"

func (m *Manager) PackFunc(info []byte) []byte {
	length := len(info)
	writer := &bytes.Buffer{}
	writer.WriteByte(methodByte)
	m.addNumber(writer, length)
	writer.Write(info)
	return writer.Bytes()
}

func (m *Manager) PackNet(list [][]byte) []byte {
	size := 0
	for _, v := range list {
		size += len(v)
	}
	writer := &bytes.Buffer{}
	m.addNumber(writer, size)
	for _, v := range list {
		writer.Write(v)
	}
	return writer.Bytes()
}
