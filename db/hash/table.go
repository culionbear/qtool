package hash

import "math"

func (m *Manager) resize() {
	oldCap, oldThr := m.cap, m.threshold
	newCap, newThr := 0, 0
	if oldCap > 0 {
		if oldCap >= maximumCapacity {
			m.threshold = math.MaxInt
			return
		} else if (oldCap << 1) < maximumCapacity && oldCap >= defaultInitialCapacity {
			newThr = oldThr << 1
			newCap = oldCap << 1
		}
	} else if oldThr > 0 {
		newCap = oldThr
	} else {
		newCap = defaultInitialCapacity
		newThr = defaultInitialCapacity * defaultLoadFactor
	}
	if newThr == 0 {
		ft := defaultLoadFactor * float32(newCap)
		if newCap < maximumCapacity && ft < float32(maximumCapacity) {
			newThr = int(ft)
		} else {
			newThr = math.MaxInt
		}
	}
	m.threshold = newThr
	newTable := make([]*link, newCap)
	//todo: table转移
	m.table = newTable
}