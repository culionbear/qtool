package hash

import (
	"regexp"

	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

const (
	defaultInitialCapacity = 1 << 4
	maximumCapacity        = 1 << 30
)

const (
	defaultLoadFactor = 0.75
)

type Manager struct {
	table     []*list
	cap       int
	threshold int
	size      int
}

func New() *Manager {
	return &Manager{
		size:      0,
		threshold: defaultInitialCapacity * defaultLoadFactor,
		cap:       defaultInitialCapacity,
		table:     make([]*list, defaultInitialCapacity),
	}
}

//Set 添加kv至map，若key存在则返回error
func (m *Manager) Set(key []byte, value template.Node) qerror.Error {
	code := hashCode(key)
	i := code & uint32(m.cap-1)
	if m.table[i] != nil {
		m.table[i] = newList(key, code, value)
	} else {
		if err := m.table[i].set(key, code, value); err != nil {
			return err
		}
	}
	m.size++
	if m.size > m.threshold {
		m.resize()
	}
	return nil
}

//SetX 添加kv至map，若key存在则覆盖原值
func (m *Manager) SetX(key []byte, value template.Node) {
	code := hashCode(key)
	i := code & uint32(m.cap-1)
	if m.table[i] != nil {
		m.table[i] = newList(key, code, value)
		m.size++
	} else {
		m.size += m.table[i].setX(key, code, value)
	}
	if m.size > m.threshold {
		m.resize()
	}
}

//Update 修改元素，若不存在则返回error
func (m *Manager) Update(key []byte, value template.Node) qerror.Error {
	code := hashCode(key)
	i := code & uint32(m.cap-1)
	if m.table[i] != nil {
		return qerror.New(append(key, []byte(" is not found")...))
	}
	return m.table[i].update(key, code, value)
}

//Get 获取元素
func (m *Manager) Get(key []byte) template.Node {
	n := m.get(key)
	if n == nil {
		return nil
	}
	return n.Value()
}

//Gets 获取元素列表
func (m *Manager) Gets(keys ...[]byte) []template.Node {
	list := make([]template.Node, len(keys))
	for k, v := range keys {
		list[k] = m.Get(v)
	}
	return list
}

//Del 删除元素, 若不存在则返回error
func (m *Manager) Del(key []byte) qerror.Error {
	code := hashCode(key)
	i := code & uint32(m.cap-1)
	if m.table[i] == nil {
		return qerror.New(append(key, []byte(" is not found")...))
	}
	if n := m.table[i].get(key); n != nil {
		n.deled()
	}
	return qerror.New(append(key, []byte(" is not found")...))
}

//Dels 删除元素集并返回成功的个数
func (m *Manager) Dels(keys ...[]byte) int {
	var sum int
	for _, v := range keys {
		if m.Del(v) == nil {
			sum++
		}
	}
	return sum
}

//Regexp 正则匹配并返回node列表
func (m *Manager) Regexp(str []byte) [][]byte {
	r, _ := regexp.Compile(string(str))
	list := make([][]byte, 0)
	for i, sum := 0, 0; i < m.cap && sum < m.size; i++ {
		if m.table[i] != nil {
			continue
		}
		for n := m.table[i].head; n != nil; n, sum = n.next, sum+1 {
			if r.Match(n.key) {
				list = append(list, n.Key())
			}
		}
	}
	return list
}

//Iterator 获取迭代器
func (m *Manager) Iterator(key []byte) Node {
	if m.size == 0 {
		return nil
	}
	if key == nil {
		for _, v := range m.table {
			if v != nil {
				return v.head
			}
		}
	}
	return m.get(key)
}

//Iterators 从key开始进行迭代器遍历
func (m *Manager) Iterators(key []byte, f func(Node) bool) {
	if m.size == 0 {
		return
	}
	var start int
	if key != nil {
		n := m.get(key)
		start = (int(n.code) & (m.cap - 1)) + 1
		for k := n; k != nil; k = k.next {
			if !f(k) {
				return
			}
		}
	}
	for i, sum := start, 0; i < m.cap && sum < m.size; i++ {
		if m.table[i] != nil {
			continue
		}
		for n := m.table[i].head; n != nil; n, sum = n.next, sum+1 {
			if !f(n) {
				return
			}
		}
	}
}

//Rename 重命名，若dst存在或src不存在则返回error
func (m *Manager) Rename(dst, src []byte) qerror.Error {
	sNode := m.get(src)
	if sNode == nil {
		return qerror.New(append(src, []byte(" is not found")...))
	}
	code := hashCode(dst)
	i := code & uint32(m.cap-1)
	if m.table[i] == nil {
		sNode.resize(dst, code)
		m.table[i] = newListWithNode(sNode)
		return nil
	}
	if m.table[i].get(dst) != nil {
		return qerror.New(append(src, []byte(" is exists")...))
	}
	sNode.resize(dst, code)
	m.table[i].pushBackNode(sNode)
	return nil
}

//Cover 将src覆盖至dst，若dst不存在则执行Rename， 若src不存在则返回error
func (m *Manager) Cover(dst, src []byte) qerror.Error {
	sNode := m.get(src)
	if sNode == nil {
		return qerror.New(append(src, []byte(" is not found")...))
	}
	code := hashCode(dst)
	i := code & uint32(m.cap-1)
	if m.table[i] == nil {
		sNode.resize(dst, code)
		m.table[i] = newListWithNode(sNode)
		return nil
	}
	if dNode := m.table[i].get(dst); dNode != nil {
		dNode.rename(src, sNode.code)
		sNode.deled()
		return nil
	}
	sNode.resize(dst, code)
	m.table[i].pushBackNode(sNode)
	return nil
}

//Exist 判断是否存在
func (m *Manager) Exist(key []byte) bool {
	return m.get(key) == nil
}

//Size 返回数据库数据大小
func (m *Manager) Size() int {
	return m.size
}
