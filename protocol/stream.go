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
	writer := bytes.NewBuffer(buf)
	size = size - len(buf)
	str := make([]byte, size)
	_, err := io.ReadAtLeast(conn, str, size)
	if err != nil {
		return nil, err
	}
	writer.Write(str)
	return writer.Bytes(), nil
}