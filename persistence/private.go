package persistence

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/culionbear/qtool/qerror"
)

const (
	fileAof	= ".aof"
	fileQdb	= ".qdb"
)

//Init persistence Manager with config path
func (m *Manager) initConfig(path string) error {
	err := NewConfig(path, &m.info)
	if err != nil {
		return err
	}
	if err := m.initAof(); err != nil {
		return err
	}
	return m.initQdb()
}

func (m *Manager) initAof() error {
	if !m.info.Aof {
		return nil
	}
	if m.info.AofPath == "" {
		m.info.AofPath = "/etc/qlite/database.aof"
	}
	if m.info.AofTimer < 15 {
		m.info.AofTimer = 15
	}
	if err := m.judgeSuffix(m.info.AofPath, fileAof); err != nil {
		return err
	}
	if err := m.judgeFile(m.info.AofPath); err != nil {
		if err = m.createFile(m.info.AofPath); err != nil {
			return err
		}
	}
	m.aofCh = make(chan *aofModule, 100000)
	m.aofCloseCh = make(chan bool)
	go m.runAof()
	return nil
}

func (m *Manager) initQdb() error {
	if m.info.Qdb == "" {
		m.info.Qdb = "/etc/qlite/database.qdb"
	}
	if m.info.QdbTimer < 120 {
		m.info.QdbTimer = 120
	}
	if err := m.judgeSuffix(m.info.Qdb, fileQdb); err != nil {
		return err
	}
	if err := m.judgeFile(m.info.Qdb); err != nil {
		if err = m.createFile(m.info.Qdb); err != nil {
			return err
		}
	}
	m.qdbCloseCh = make(chan bool)
	go m.runQdb()
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
			buf, err := m.serialize(msg)
			if err != nil {
				continue
			}
			writer.Write(buf)
		case <- m.aofCloseCh:
			_ = writer.Flush()
			return
		case <- timer.C:
			_ = writer.Flush()
		}
	}
}

func (m *Manager) runQdb() {
	defer func() {
		m.closeCh <- true
	}()

	timer := time.NewTicker(time.Duration(m.info.QdbTimer) * time.Second)

	for {
		select {
		case <- m.qdbCloseCh:
			return
		case <- timer.C:
			
		}
	}
}

func (m *Manager) judgeSuffix(path, suffix string) error {
	if strings.LastIndex(path, suffix) == -1 {
		return qerror.NewString("suffix error")
	}
	return nil
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

func (m *Manager) serialize(msg *aofModule) ([]byte, error) {
	switch msg.Cmd {
	case CmdSet:
		return msg.getSetModule()
	case CmdSetX:
		return msg.getSetXModule()
	case CmdUpdate:
		return msg.getUpdateModule()
	case CmdDel:
		return msg.getDelModule()
	case CmdDels:
		return msg.getDelsModule()
	case CmdRename:
		return msg.getRenameModule()
	case CmdCover:
		return msg.getCoverModule()
	default:
		return nil, qerror.New([]byte("cmd is not found"))
	}
}