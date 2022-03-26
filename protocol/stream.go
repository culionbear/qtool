package protocol

import (
	"bytes"
	"io"
	"net"
)

func (m *Manager) PackSize(buf []byte) (int, int, bool) {
	for k, v := range buf {
		if v == separator {
			m.init()
			return m.size, k, true
		}
		m.size += m.mod * int(v)
		m.mod *= modValue
	}
	return 0, 0, false
}

func (m *Manager) Read(conn net.Conn, size int, buf []byte) ([]byte, error) {
	size = size - len(buf)
	if size <= 0 {
		return buf[: size], nil
	}
	writer := bytes.NewBuffer(buf)
	str := make([]byte, size)
	_, err := io.ReadAtLeast(conn, str, size)
	if err != nil {
		return nil, err
	}
	writer.Write(str)
	return writer.Bytes(), nil
}

func (m *Manager) Write(buf []byte) []byte {
	size := len(buf)
	writer := &bytes.Buffer{}
	m.addNumber(writer, size)
	writer.Write(buf)
	return writer.Bytes()
}
