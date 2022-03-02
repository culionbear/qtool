package hash

import (
	"github.com/culionbear/qtool/db/iterator"
	"github.com/culionbear/qtool/qerror"
	"github.com/culionbear/qtool/template"
)

const(
	defaultInitialCapacity	= 1 << 4
	maximumCapacity			= 1 << 30
)

const(
	defaultLoadFactor		= 0.75
)

type Manager struct {
	table		[]*link
	cap			int
	threshold	int
	size		int
}

func New() *Manager {
	return &Manager{
		size: 0,
		threshold: defaultInitialCapacity * defaultLoadFactor,
		cap: defaultInitialCapacity,
		table: make([]*link, defaultInitialCapacity),
	}
}

//Set 添加kv至map，若key存在则返回error
func (m *Manager) Set(key []byte, value template.Node) qerror.Error {
	return nil
}

//SetX 添加kv至map，若key存在则覆盖原值
func (m *Manager) SetX(key []byte, value template.Node) {
	code := hashCode(key)
	i := code & uint32(m.cap - 1)
	if m.table[i] != nil {
		m.table[i] = newLink(code, key, value)
		m.size ++
	} else {

	}
	if m.size > m.threshold {
		m.resize()
	}
}

//Get 获取元素
func (m *Manager) Get(key []byte) (template.Node, qerror.Error) {
	return nil, nil
}

//Del 删除元素并返回成功删除个数
func (m *Manager) Del(key... []byte) int {
	return 0
}

//Regexp 正则匹配并返回node列表
func (m *Manager) Regexp(str []byte) ([]template.Node, qerror.Error) {
	return nil, nil
}

//Iterator 获取迭代器
func (m *Manager) Iterator(key []byte) (iterator.Node, qerror.Error) {
	return nil, nil
}

//Rename 重命名，若dst存在则返回error
func (m *Manager) Rename(dst, src []byte) qerror.Error {
	return nil
}

//Cover 将src覆盖至dst，若dst不存在则执行Rename
func (m *Manager) Cover(dst, src []byte) qerror.Error{
	return nil
}

func (m *Manager) Exist(key []byte) bool {
	return false
}

//Size 返回数据库数据大小
func (m *Manager) Size() int {
	return m.size
}