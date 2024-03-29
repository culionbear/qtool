package hash

import (
	"bytes"
)

func compare(a, b []byte) bool {
	return bytes.Compare(a, b) == 0
}

func hashCode(buf []byte) uint32 {
	var sum uint32
	for _, v := range buf {
		sum = sum<<5 - sum + uint32(v)
	}
	return sum ^ (sum >> 16)
}
