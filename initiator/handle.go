package initiator

var m *Manager

// func init() {
// 	m = NewManager()
// 	go m.Run()
// }

func Close() {
	m.Close()
}

func Set(msg *Module) {
	m.Set(msg)
}
