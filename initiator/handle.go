package initiator

var m *manager

func init() {
	m = newManager()
	go m.run()
}

func Close() {
	m.close()
}

func Set(msg *Module) {
	m.set(msg)
}
