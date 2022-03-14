package persistence

import (
	"bufio"
	"os"
	"time"

	"github.com/culionbear/qtool/db/hash"
)

//Manager persistence
type Manager struct {
	info		Config
	aofCh		chan AofModule
	aofCloseCh	chan bool
	closeCh		chan bool
}

//New Manager
func New() *Manager {
	return &Manager{}
}

//Init persistence Manager with config path
func (m *Manager) Init(path string) error {
	err := NewConfig(path, &m.info)
	if err != nil {
		return err
	}
	if m.info.Aof {
		if m.info.AofPath == "" {
			m.info.AofPath = "/etc/qlite/database.aof"
		}
		if err = m.judgeFile(m.info.AofPath); err != nil {
			if err = m.createFile(m.info.AofPath); err != nil {
				return err
			}
		}
		m.aofCh = make(chan AofModule, 100000)
		m.aofCloseCh = make(chan bool)
		go m.runAof()
	}
	if err = m.judgeFile(m.info.QdbPath); err != nil {
		if err = m.createFile(m.info.QdbPath); err != nil {
			return err
		}
	}
	return nil
}

//Close persistence Manager
func (m *Manager) Close() {
	if m.info.Aof {
		m.aofCloseCh <- true
		<- m.closeCh
	}

}

func (m *Manager) QdbSave(values []hash.Node) error {
	fp, err := os.OpenFile("./qdb.duplicate", os.O_WRONLY | os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()
	for _, v := range values {
		_ = v
	}

	return nil
}

func (m *Manager) runAof() {
	fp, err := os.OpenFile(m.info.AofPath, os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer fp.Close()
	defer func() {
		m.closeCh <- true
	}()

	writer := bufio.NewWriter(fp)

	timer := time.NewTicker(time.Duration(m.info.AofTimer) * time.Second)

	for {
		select {
		case msg := <- m.aofCh:
			_ = msg
		case <- m.aofCloseCh:
			_ = writer.Flush()
			return
		case <- timer.C:
			_ = writer.Flush()
		}
	}
}

func (m *Manager) judgeFile(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	return nil
}

func (m *Manager) createFile(path string) error {
	_, err := os.Create(path)
	return err
}