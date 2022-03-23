package initiator

type manager struct {
	ch		chan *Module
	chClose	chan bool
}

func newManager() *manager {
	return &manager{
		ch: make(chan *Module, 100000),
		chClose: make(chan bool),
	}
}

func (m *manager) close() {
	m.chClose <- true
}

func (m *manager) run() {
	for {
		select {
		case msg := <- m.ch:
			if len(msg.list) == 0 {

			}
		case <- m.chClose:
			return
		}
	}
}

func (m *manager) set(msg *Module) {
	m.ch <- msg
}