package protocol

import (
	"encoding/binary"
	"math"

	"github.com/culionbear/qtool/ds/queue"
	"github.com/culionbear/qtool/qerror"
)

func (m *Manager) Unpack(buf []byte) (*queue.Manager[any], qerror.Error) {
	var err qerror.Error
	q, size := queue.New[any](), len(buf)
	for i := 0; i < size; {
		i, err = m.unpack(q, i, size, buf)
		if err != nil {
			return nil, err
		}
	}
	return q, nil
}

func (m *Manager) readSize(i, size int, buf []byte) (int, int, qerror.Error) {
	sum, mod := 0, 1
	for j := i; j < size; j ++ {
		if buf[j] == separator {
			return sum, j + 1, nil
		}
		sum += mod * int(buf[j])
		mod *= modValue
	}
	return 0, 0, qerror.NewString("package length is error")
}

func (m *Manager) copyBuf(buf []byte) []byte {
	dst := make([]byte, len(buf))
	copy(dst, buf)
	return dst
}

func (m *Manager) unpack(q *queue.Manager[any], i, size int, buf []byte) (int, qerror.Error) {
	switch buf[i] {
	case trueByte:
		q.Push(true)
		return i + 1, nil
	case falseByte:
		q.Push(false)
		return i + 1, nil
	case numberByte:
		sum, j, err := m.readSize(i + 1, size, buf)
		if err != nil {
			return 0, err
		}
		q.Push(sum)
		return j, nil
	case floatByte:
		if i + 9 > size {
			return 0, qerror.NewString("package length is error")
		}
		q.Push(
			math.Float64frombits(
				binary.LittleEndian.Uint64(buf[i + 1: i + 9]),
			),
		)
		return i + 9, nil
	case stringByte:
		sum, j, err := m.readSize(i + 1, size, buf)
		if err != nil {
			return 0, err
		}
		if j + sum > size {
			return 0, qerror.NewString("package length is error")
		}
		q.Push(m.copyBuf(buf[j: j + sum]))
		return j + sum, nil
	case errorByte:
		sum, j, err := m.readSize(i + 1, size, buf)
		if err != nil {
			return 0, err
		}
		if j + sum > size {
			return 0, qerror.NewString("package length is error")
		}
		q.Push(qerror.Copy(buf[j: j + sum]))
		return j + sum, nil
	case listByte:
		ql := queue.New[any]()
		var err qerror.Error
		for buf[i] == listByte {
			i, err = m.unpack(ql, i + 1, size, buf)
			if err != nil {
				return 0, err
			}
		}
		return i, nil
	default:
		return 0, qerror.NewString("unknown type")
	}
}
