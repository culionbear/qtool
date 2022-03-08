package hash

import (
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

const(
	defaultInitialCapacity	= 1 << 4
	maximumCapacity			= 1 << 30
)

const(
	treeifyThreshold		= 8
	unTreeifyThreshold		= 6
	minTreeifyCapacity		= 64
)

const(
	defaultLoadFactor		= 0.75
)

type Manager struct {
	table		[]Node
	cap			int
	threshold	int
	size		int
}

func New() *Manager {
	return &Manager{
		size: 0,
		threshold: defaultInitialCapacity * defaultLoadFactor,
		cap: defaultInitialCapacity,
		table: make([]Node, defaultInitialCapacity),
	}
}

//Set 添加kv至map，若key存在则返回error
func (m *Manager) Set(key []byte, value template.Node) qerror.Error {
	code := hashCode(key)
	i := code & uint32(m.cap - 1)
	if m.table[i] != nil {
		m.table[i] = newListNode(key, code, value)
	} else {
		if m.table[i].Hex() == code && compare(m.table[i].unsafeGetKey(), key) {
			return qerror.New(append(key, []byte(" is exists")...))
		} else if n, ok := m.table[i].(*treeNode); ok {//is rbt
			_ = n
		} else {
			h := m.table[i].(*listNode)
			for n, size := h.next, 0; ; n, size = n.next, size + 1 {
				if n.code == code && compare(n.key, key) {
					return qerror.New(append(key, []byte(" is exists")...))
				}
				if n.next == nil {
					n.next = newListNode(key, code, value)
					if size >= treeifyThreshold - 1 {
						m.treeifyBin(h)
					}
					break
				}
			}
		}
	}
	m.size ++
	if m.size > m.threshold {
		m.resize()
	}
	return nil
}

//SetX 添加kv至map，若key存在则覆盖原值
func (m *Manager) SetX(key []byte, value template.Node) {
	code := hashCode(key)
	i := code & uint32(m.cap - 1)
	if m.table[i] != nil {
		m.table[i] = newListNode(key, code, value)
		m.size ++
	} else {
		if m.table[i].Hex() == code && compare(m.table[i].unsafeGetKey(), key) {
			m.table[i].SetValue(value)
		} else if n, ok := m.table[i].(*treeNode); ok {//is rbt
			_ = n
		} else {
			h := m.table[i].(*listNode)
			for n, size := h.next, 0; ; n, size = n.next, size + 1 {
				if n.code == code && compare(n.key, key) {
					n.value = value
					break
				}
				if n.next == nil {
					n.next = newListNode(key, code, value)
					m.size ++
					if size >= treeifyThreshold {
						m.treeifyBin(h)
					}
					break
				}
			}
		}
	}
	if m.size > m.threshold {
		m.resize()
	}
}

//Update 修改元素，若不存在则返回error
func (m *Manager) Update(key []byte, value template.Node) qerror.Error {
	code := hashCode(key)
	i := code & uint32(m.cap - 1)
	if m.table[i] != nil {
		return qerror.New(append(key, []byte(" is not found")...))
	}
	if m.table[i].Hex() == code && compare(m.table[i].unsafeGetKey(), key) {
		m.table[i].SetValue(value)
	} else if n, ok := m.table[i].(*treeNode); ok {//is rbt
		_ = n
	} else {
		h := m.table[i].(*listNode)
		for n, size := h.next, 0; ; n, size = n.next, size + 1 {
			if n.code == code && compare(n.key, key) {
				n.value = value
				break
			}
			if n.next == nil {
				return qerror.New(append(key, []byte(" is not found")...))
			}
		}
	}
	return nil
}

//Get 获取元素
func (m *Manager) Get(key []byte) template.Node {
	code := hashCode(key)
	i := code & uint32(m.cap - 1)
	if m.table[i] == nil {
		return nil
	}
	if compare(m.table[i].unsafeGetKey(), key) {
		return m.table[i].Value()
	} else {
		for n := m.table[i].Next(); n != nil; n = n.Next() {
			if compare(n.unsafeGetKey(), key) {
				return n.Value()
			}
		}
	}
	return nil
}

//Gets 获取元素列表
func (m *Manager) Gets(keys... []byte) []template.Node {
	list := make([]template.Node, len(keys))
	for k, v := range keys {
		list[k] = m.Get(v)
	}
	return list
}

//Del 删除元素并返回成功删除个数
func (m *Manager) Del(keys... []byte) int {
	return 0
}

//Regexp 正则匹配并返回node列表
func (m *Manager) Regexp(str []byte) ([]template.Node, qerror.Error) {
	return nil, nil
}

//Iterator 获取迭代器
func (m *Manager) Iterator(key []byte) (Node, qerror.Error) {
	return nil, nil
}

//Iterators 从key开始进行迭代器遍历
func (m *Manager) Iterators(key []byte, f func(Node) bool) qerror.Error {
	return nil
}

//Rename 重命名，若dst存在则返回error
func (m *Manager) Rename(dst, src []byte) qerror.Error {
	return nil
}

//Cover 将src覆盖至dst，若dst不存在则执行Rename
func (m *Manager) Cover(dst, src []byte) qerror.Error{
	return nil
}

//Exist 判断是否存在
func (m *Manager) Exist(key []byte) bool {
	code := hashCode(key)
	i := code & uint32(m.cap - 1)
	if m.table[i] == nil {
		return false
	}
	if compare(m.table[i].unsafeGetKey(), key) {
		return true
	} else if n, ok := m.table[i].(*treeNode); ok {//is rbt
		_ = n
	} else {
		h := m.table[i].(*listNode)
		for n := h.next; n != nil; n = n.next {
			if compare(n.key, key) {
				return true
			}
		}
	}
	return false
}

//Size 返回数据库数据大小
func (m *Manager) Size() int {
	return m.size
}