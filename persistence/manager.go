package persistence

//Manager persistence
type Manager struct {
	info		Config
	aofCh		chan *aofModule
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

//Close persistence Manager
func (m *Manager) Close() {
	m.aofCloseCh <- true
	<- m.closeCh
}

//Save logs in local
func (m *Manager) Save(cmd uint8, args []interface{}) {
	m.aofCh <- newAofModule(cmd, args)
}