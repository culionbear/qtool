package hash

import (
	"fmt"
	"strconv"
	"testing"
)

type testNode struct {
}

func (m *testNode) Type() []byte {
	return []byte("test")
}

func (m *testNode) Serialize() []byte {
	return []byte{}
}

func (m *testNode) Deseriallize([]byte) {

}

func BenchmarkHash(t *testing.B) {
	m := New()
	l := make([][]byte, 0)
	sum := 100000
	for i := 0; i < sum; i++ {
		l = append(l, []byte(strconv.Itoa(i)))
	}
	n := new(testNode)
	t.ResetTimer()
	for i := 0; i < sum; i++ {
		m.Set(l[i], n)
		// if err != nil {
		// 	t.Log(i, err)
		// }
		// t.Logf("%6d %6d %6d", m.size, m.cap, m.threshold)
	}
	t.StopTimer()
	t.StartTimer()
	for i := 0; i < sum; i++ {
		err := m.Del(l[i])
		if err != nil {
			t.Log(i, err)
		}
	}
}

func TestHash(t *testing.T) {
	m := New()
	l := make([][]byte, 0)
	sum := 100000
	for i := 0; i < sum; i++ {
		l = append(l, []byte(strconv.Itoa(i)))
	}
	n := new(testNode)
	//th := 128
	for i := 0; i < sum; i++ {
		err := m.Set(l[i], n)
		if err != nil {
			t.Log(i, err)
		}
	}
	t.Log(m.Dels(l...))
	//PrintTable(t, m)
	// for i := 0; i < sum; i ++ {
	// 	n := m.Get(l[i])
	// 	if n != nil {
	// 		t.Log("get", i)
	// 	}
	// }

	a := 0
	m.Iterators(nil, func(Node) bool {
		a++
		return true
	})
	t.Log(a)
	//t.Log(string(m.Iterator(l[5000]).Key()))
}

func PrintTable(t *testing.T, m *Manager) {
	l := make([]interface{}, 0)
	for k, v := range m.table {
		l = append(l, fmt.Sprintf("table[%d]:", k))
		if v != nil {
			for n := v.head; n != nil; n = n.next {
				if n != v.head {
					l = append(l, ",")
				}
				l = append(l, string(n.key))
			}
		}
		l = append(l, "\n")
	}
	t.Log(l...)
	t.Log("------------------------------------------------------------------------------")
}
