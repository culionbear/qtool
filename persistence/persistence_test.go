package persistence

import "testing"

func TestXxx(t *testing.T) {
	msg := []byte("string")
	t.Log(pack(msg), len(msg))
}