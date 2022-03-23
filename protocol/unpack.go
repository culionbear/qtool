package protocol

import (
	"encoding/binary"
	"math"

	"github.com/culionbear/qtool/ds/queue"
	"github.com/culionbear/qtool/qerror"
)

func (m *Manager) Unpack(buf []byte) (*queue.Manager, qerror.Error) {
	q, size := queue.New(), len(buf)
	for i := 0; i < size; {
		switch buf[i] << cmdValue >> cmdValue {
		case trueByte:
			q.Push(true)
			i ++
		case falseByte:
			q.Push(false)
			i ++
		case numberByte:
			sum, j, err := m.readSize(i + 1, size, buf)
			if err != nil {
				return nil, err
			}
			q.Push(sum)
			i = j
		case floatByte:
			if i + 9 > size {
				return nil, qerror.NewString("package length is error")
			}
			q.Push(
				math.Float64frombits(
					binary.LittleEndian.Uint64(buf[i + 1: i + 9]),
				),
			)
			i += 9
		case stringByte:
			sum, j, err := m.readSize(i + 1, size, buf)
			if err != nil {
				return nil, err
			}
			if j + sum > size {
				return nil, qerror.NewString("package length is error")
			}
			q.Push(m.copyBuf(buf[j: j + sum]))
			i = j + sum
		case errorByte:
			sum, j, err := m.readSize(i + 1, size, buf)
			if err != nil {
				return nil, err
			}
			if j + sum > size {
				return nil, qerror.NewString("package length is error")
			}
			q.Push(qerror.Copy(buf[j: j + sum]))
			i = j + sum
		case listByte:
		default:
			return nil, qerror.NewString("unknown type")
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