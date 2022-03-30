package protocol

import (
	"bytes"
	"io"
	"net"
)

func (m *Manager) PackSize(buf []byte) (int, int, bool) {
	for k, v := range buf {
		if v == separator {
			defer m.init()
			return m.size, k + 1, true
		}
		m.size += m.mod * int(v)
		m.mod *= modValue
	}
	return 0, 0, false
}

func (m *Manager) Read(conn net.Conn, size int, buf []byte) ([]byte, error) {
	if size-len(buf) <= 0 {
		return buf[:size], nil
	}
	size = size - len(buf)
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
	writer := &bytes.Buffer{}
	m.addNumber(writer, len(buf))
	writer.Write(buf)
	return writer.Bytes()
}
