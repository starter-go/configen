package gocode

type Import struct {
	Alias    string
	FullName string
}

// TypeStruct 表示一个结构体类型
type TypeStruct struct {
	IsComponent bool            // 是否是 starter 组件
	ComAtts     *ConfigenParams // starter 组件属性
	ComID       string
	ComClass    string
	ComAlias    string
	ComScope    string

	Name   string   // 结构体名称
	Fields FieldSet // 字段集合
}

type Field struct {
	Name      string
	Type      ComplexType
	Injection string
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
