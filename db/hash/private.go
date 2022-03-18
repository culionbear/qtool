package hash

func (m *Manager) get(key []byte) *node {
	i := hashCode(key) & uint32(m.cap-1)
	if m.table[i] == nil {
		return nil
	}
	return m.table[i].get(key)
}

func (m *Manager) del(n *node) {
	i := n.code & uint32(m.cap-1)
	if m.table[i].delNode(n) {
		m.table[i] = nil
	}
}
