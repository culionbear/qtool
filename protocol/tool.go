package protocol

import (
	"reflect"
	"unsafe"
)

func (m *Manager) s2b(str string) (buf []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return
}
