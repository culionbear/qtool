package persistence

//Manager persistence
type Manager struct {
	info		Config
	aofCh		chan *module
	aofCloseCh	chan bool
	closeCh		chan bool
}

//New Manager
func New(path string) (*Manager, error) {
	m := &Manager{}
	return m, m.initConfig(path)
}

//NewWithConfig to Manager
func NewWithConfig(c Config) (*Manager, error) {
	m := &Manager{
		info: c,
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
	<- m.closeCh
}

//Save logs in local
func (m *Manager) Save(cmd uint8, args []interface{}) {
	m.aofCh <- newModule(cmd, args)
}

//Fetch local in table
func (m *Manager) Fetch(f func([]byte, [][]byte) error) error {
	buf, err := m.readAll()
	if err != nil {
		return err
	}
	//TODO:恢复数据
	_ = buf
	return nil
}
