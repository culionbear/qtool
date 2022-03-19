package persistence

import (
	"github.com/culionbear/qtool/logs"
	"github.com/culionbear/qtool/qerror"
)

//Manager persistence
type Manager struct {
	info       Config
	aofCh      chan *module
	aofCloseCh chan bool
	closeCh    chan bool
	cmdTable   *CmdOpt[moduleFunc]
}

//New Manager
func New(path string) (*Manager, error) {
	m := &Manager{}
	m.cmdTable = &CmdOpt[moduleFunc]{
		CmdSet:    m.getSetModule,
		CmdSetX:   m.getSetXModule,
		CmdUpdate: m.getUpdateModule,
		CmdDel:    m.getDelModule,
		CmdDels:   m.getDelsModule,
		CmdRename: m.getRenameModule,
		CmdCover:  m.getCoverModule,
	}
	return m, m.initConfig(path)
}

//NewWithConfig to Manager
func NewWithConfig(c Config) (*Manager, error) {
	m := &Manager{
		info: c,
	}
	m.cmdTable = &CmdOpt[moduleFunc]{
		CmdSet:    m.getSetModule,
		CmdSetX:   m.getSetXModule,
		CmdUpdate: m.getUpdateModule,
		CmdDel:    m.getDelModule,
		CmdDels:   m.getDelsModule,
		CmdRename: m.getRenameModule,
		CmdCover:  m.getCoverModule,
	}
	return m, m.initAof()
}

//Run aof gorountie
func (m *Manager) Run() {
	go m.runAof()
}

//Close persistence Manager
func (m *Manager) Close() {
	m.aofCloseCh <- true
	<-m.closeCh
}

//Save logs in local
func (m *Manager) Save(cmd uint8, args []any) {
	m.aofCh <- newModule(cmd, args)
}

//Fetch local in table
func (m *Manager) Fetch(f func(uint8, [][]byte) error) error {
	buf, err := m.readAll()
	if err != nil {
		return err
	}
	length, success, all := uint64(len(buf)), 0, 0
	for i := uint64(0); i < length; all++ {
		size, err := getPackageLength(buf)
		if err != nil {
			return err
		}
		i += headLength
		if i+size > length {
			return qerror.NewString("bytes size is error")
		}
		var cmd uint8 = buf[i]
		list := make([][]byte, 0)
		sum := i + size
		for j := i + 1; j < sum; {
			pSize, err := getPackageLength(buf[j:sum])
			if err != nil {
				return err
			}
			j += headLength
			if j+pSize > sum {
				return qerror.NewString("package size is error")
			}
			list = append(list, m.copyAll(buf[j:j+pSize]))
			j += pSize
		}
		if err = f(cmd, list); err != nil {
			logs.PrintError(err)
		} else {
			success++
		}
		i = sum
	}
	logs.PrintInfo("fetch complete.success:", success, ";all:", all)
	return nil
}
