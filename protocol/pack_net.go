package protocol

import "bytes"

func (m *Manager) PackFunc(info []byte) []byte {
	length := len(info)
	writer := &bytes.Buffer{}
	writer.WriteByte(funcByte)
	m.addNumber(writer, length)
	writer.WriteByte(separator)
	writer.Write(info)
	return writer.Bytes()
}

func (m *Manager) PackNet(list [][]byte) []byte {
	length := len(list)
	writer := &bytes.Buffer{}
	for i := 0; i < length; i ++ { 
		list[i][0] |= chByte
		writer.Write(list[i])
	}
	return writer.Bytes()
}
