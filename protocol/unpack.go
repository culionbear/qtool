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
	case byteByte:
		if i + 2 > size {
			return 0, qerror.NewString("package length is error")
		}
		q.Push(buf[i + 1])
		return i + 2, nil
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
	default:
		if buf[i] & listByte != listByte {
			return 0, qerror.NewString("unknown type")
		}
		return m.unpackList(q, i, size, buf)
	}
}

func (m *Manager) unpackList(q *queue.Manager[any], i, size int, buf []byte) (int, qerror.Error) {
	sum, j, err := m.readSize(i + 1, size, buf)
	if err != nil {
		return 0, err
	}
	switch buf[i] << separatorByte >> separatorByte {
	case stringByte:
		list := make([]string, sum, sum)
		for k := 0; k < sum; k ++ {
			lSize, next, err := m.readSize(j, size, buf)
			if err != nil {
				return 0, nil
			}
			if next + lSize > size {
				return 0, qerror.NewString("package length is error")
			}
			list[k] = string(buf[next: next + lSize])
			j = next + lSize
		}
		q.Push(list)
		return j, nil
	case errorByte:
		list := make([]qerror.Error, sum, sum)
		for k := 0; k < sum; k ++ {
			lSize, next, err := m.readSize(j, size, buf)
			if err != nil {
				return 0, nil
			}
			if next + lSize > size {
				return 0, qerror.NewString("package length is error")
			}
			list[k] = qerror.Copy(buf[next: next + lSize])
			j = next + lSize
		}
		q.Push(list)
		return j, nil
	case trueByte:
		if j + sum > size {
			return 0, qerror.NewString("package length is error")
		}
		list := make([]bool, sum, sum)
		for k := 0; k < sum; k ++ {
			list[k] = buf[j + k] == trueByte
		}
		q.Push(list)
		return j + sum, nil
	case numberByte:
		list := make([]int, sum, sum)
		for k := 0; k < sum; k ++ {
			num, next, err := m.readSize(j, size, buf)
			if err != nil {
				return 0, nil
			}
			list[k] = num
			j = next
		}
		q.Push(list)
		return j, nil
	case floatByte:
		if j + (sum * 8) > size {
			return 0, qerror.NewString("package length is error")
		}
		list := make([]float64, sum, sum)
		for k := 0; k < sum; k ++ {
			list[k] = math.Float64frombits(
				binary.LittleEndian.Uint64(buf[j + (8 * k): j + 8 + (8 * k)]),
			)
		}
		q.Push(list)
		return j + (sum * 8), nil
	default:
		return 0, qerror.NewString("unknown type")
	}
}
