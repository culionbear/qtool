package persistence

import (
	"os"
)

//Manager persistence
type Manager struct {
	info		Config
	aofCh		chan *aofModule
	aofCloseCh	chan bool
	qdbCloseCh	chan bool
	closeCh		chan bool
}

//New Manager
func New(path string) (*Manager, error) {
	m := &Manager{}
	return m, m.initConfig(path)
}

//Close persistence Manager
func (m *Manager) Close() {
	if m.info.Aof {
		m.aofCloseCh <- true
		<- m.closeCh
	}
	m.qdbCloseCh <- true
	<- m.closeCh
}

//QdbSave save binary data in local
func (m *Manager) QdbSave(nodes []*QdbModule) error {
	fp, err := os.OpenFile("./qdb.duplicate", os.O_WRONLY | os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()
	
	return nil
}

//AofSave save logs in local
func (m *Manager) AofSave(cmd uint8, args []interface{}) {
	if !m.info.Aof {
		return
	}
	m.aofCh <- newAofModule(cmd, args)
}