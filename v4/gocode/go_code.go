package gocode

import (
	"crypto/md5"

	"github.com/starter-go/base/lang"
)

// Import 表示导入当前源文件的一个包
type Import struct {
	Alias    string
	FullName string
	HexName  string
}

// ComputeHexName 计算字段 'HexName' 的值
func (inst *Import) ComputeHexName() {
	plain := inst.FullName
	sum := md5.Sum([]byte(plain))
	hex := lang.HexFromBytes(sum[:])
	inst.HexName = hex.String()
}

////////////////////////////////////////////////////////////////////////////////

// TypeStruct 表示一个结构体类型
type TypeStruct struct {
	IsComponent bool            // 是否是 starter 组件
	ComAtts     *ConfigenParams // starter 组件属性
	ComID       string
	ComClass    string
	ComAlias    string
	ComScope    string

	Name   string            // 结构体名称
	Fields FieldSet          // 字段集合
	As     ImplementationSet // 这里用 Field 结构来表示实现的各个接口
}

// Field 表示 struct 中的一个需要注入的字段
type Field struct {
	Name      string
	Type      ComplexType
	Injection string
}

////////////////////////////////////////////////////////////////////////////////

// Implementation 表示组件实现的一个接口
type Implementation struct {
	Field
}

// ImplementationSet 表示组件实现的一组接口
type ImplementationSet struct {
	list []*Implementation
}

// Add ...
func (inst *ImplementationSet) Add(item *Implementation) {
	if item == nil || inst == nil {
		return
	}
	inst.list = append(inst.list, item)
}

////////////////////////////////////////////////////////////////////////////////

// ImportSet 是若干 Import 的集合
type ImportSet struct {
	list  []*Import
	table map[string]*Import
}

// Add ...
func (inst *ImportSet) Add(items ...*Import) {
	table := inst.table
	if table == nil {
		table = make(map[string]*Import)
		inst.table = table
	}
	for _, item := range items {
		if item == nil {
			continue
		}
		table[item.Alias] = item
		table[item.FullName] = item
		inst.list = append(inst.list, item)
	}
}

// Find 通过别名或者全名查找
func (inst *ImportSet) Find(name string) *Import {
	table := inst.table
	if table == nil {
		return nil
	}
	return table[name]
}

////////////////////////////////////////////////////////////////////////////////

// TypeStructSet 是若干 TypeStruct 的集合
type TypeStructSet struct {
	list []*TypeStruct
}

// Add ...
func (inst *TypeStructSet) Add(items ...*TypeStruct) {
	inst.list = append(inst.list, items...)
}

////////////////////////////////////////////////////////////////////////////////

// FieldSet 是若干 Field 的集合
type FieldSet struct {
	list []*Field
}

// Add ...
func (inst *FieldSet) Add(items ...*Field) {
	inst.list = append(inst.list, items...)
}

////////////////////////////////////////////////////////////////////////////////
