package hash

import (
	"math"
)

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
	newTable := make([]Node, newCap)
	for i := 0; i < oldCap; i ++ {
		if m.table[i] == nil {
			continue
		}
		if m.table[i].Next() == nil {
			newTable[m.table[i].Hex() & uint32(newCap - 1)] = m.table[i]
		} else if n, ok := m.table[i].(*treeNode); ok { //is rbt
			_ = n
		} else {
			h := m.table[i].(*listNode)
			var loHead, loTail *listNode
			var hiHead, hiTail *listNode
			for n := h.next; n != nil; n = n.next {
				if n.code & uint32(oldCap) == 0 {
					if loTail == nil {
						loHead = n
					} else {
						loTail.next = n
					}
					loTail = n
				} else {
					if hiTail == nil {
						hiHead = n
					} else {
						hiTail.next = n
					}
					hiTail = n
				}
			}
			if loTail != nil {
				newTable[i] = loHead
			}
			if hiTail != nil {
				newTable[i + oldCap] = hiHead
			}
		}
	}
	m.cap = newCap
	m.table = newTable
}

func (m *Manager) treeifyBin(e *listNode) {
	if m.size < minTreeifyCapacity {
		m.resize()
		return
	}

}