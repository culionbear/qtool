package persistence

import (
	"bufio"
	"io"
	"os"
	"strings"
	"time"

	"github.com/culionbear/qtool/logs"
	"github.com/culionbear/qtool/qerror"
)

const (
	fileAof = ".aof"
	fileQdb = ".qdb"
)

//Init persistence Manager with config path
func (m *Manager) initConfig(path string) error {
	err := NewConfig(path, &m.info)
	if err != nil {
		return err
	}
	return m.initAof()
}

func (m *Manager) initAof() error {
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
	m.aofCh = make(chan *module, 100000)
	m.aofCloseCh = make(chan bool)
	return nil
}

func (m *Manager) runAof() {
	fp, err := os.OpenFile(m.info.AofPath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logs.PrintError(err)
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
		case msg := <-m.aofCh:
			buf, err := m.serialize(msg)
			if err != nil {
				logs.PrintError(err)
				continue
			}
			writer.Write(buf)
		case <-m.aofCloseCh:
			_ = writer.Flush()
			return
		case <-timer.C:
			_ = writer.Flush()
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

func (m *Manager) serialize(msg *module) ([]byte, error) {
	if f := m.cmdTable.Get(msg.Cmd); f != nil {
		return f(msg)
	}
	return nil, qerror.NewString("cmd is not found")
}

func (m *Manager) readAll() ([]byte, error) {
	fp, err := os.Open(m.info.AofPath)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	return io.ReadAll(fp)
}

func (m *Manager) copyAll(buf []byte) []byte {
	b := make([]byte, len(buf))
	copy(b, buf)
	return b
}
