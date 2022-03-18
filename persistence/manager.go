package persistence

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
func (m *Manager) Save(cmd uint8, args []interface{}) {
	m.aofCh <- newModule(cmd, args)
}

//Fetch local in table
func (m *Manager) Fetch(f func(uint8, [][]byte) error) error {
	buf, err := m.readAll()
	if err != nil {
		return err
	}
	//TODO:恢复数据
	_ = buf
	return nil
}
