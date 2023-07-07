package codetemplate

import (
	"strings"
	"text/template"

	_ "embed"
)

//go:embed the_type_struct.template
var theTemplateTextForTypeStruct string

// SimpleTemplate 简单配置-代码模板
type SimpleTemplate struct {
	templateTS *template.Template
}

// MakeTypeStruct 创建 'type Name struct {}' 代码片段
func (inst *SimpleTemplate) MakeTypeStruct(ts *TypeStruct) (string, error) {
	t := inst.templateTS
	if t == nil {
		name := "the_type_struct.template"
		raw := theTemplateTextForTypeStruct
		t2, err := template.New(name).Parse(raw)
		if err != nil {
			return "", err
		}
		t = t2
		inst.templateTS = t2
	}
	builder := &strings.Builder{}
	err := t.Execute(builder, ts)
	if err != nil {
		return "", err
	}
	str := builder.String()
	return str, nil
}

////////////////////////////////////////////////////////////////////////////////

// TypeStruct 是 MakeTypeStruct 的参数
type TypeStruct struct {
	InstanceType         string
	InstancePackageName  string
	InstancePackageAlias string

	FactoryType         string
	FactoryPackageName  string
	FactoryPackageAlias string

	ComID    string
	ComClass string
	ComAlias string
	ComScope string

	InjectFieldList    string
	InjectFieldGetters string
}
