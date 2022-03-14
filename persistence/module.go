package persistence

import (
	"github.com/culionbear/qtool/db/hash"
)

//qdbModule struct in binary file
type qdbModule struct {
	key		[]byte
	class	[]byte
	buf		[]byte
}

//newQdbModule return Module point with hash.Node
func newQdbModule(value hash.Node) *qdbModule {
	return &qdbModule {
		key:	value.Key(),
		class:	value.Value().Type(),
		buf:	value.Value().Serialize(),
	}
}

//AofModule struct in log file
type AofModule struct {
	Cmd		uint8
	Buf		[]byte
}

func NewAofModule(cmd uint8, buf []byte) *AofModule {
	return &AofModule{
		Cmd: cmd,
		Buf: buf,
	}
}