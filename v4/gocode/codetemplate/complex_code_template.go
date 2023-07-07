package codetemplate

import (
	_ "embed"
	"strings"
	"text/template"
)

//go:embed the_auto_config.template
var theTemplateTextForAutoConfig string

////////////////////////////////////////////////////////////////////////////////

// AutoConfig 是函数 MakeAutoConfigFunc 的参数
type AutoConfig struct {
	PackageSimpleName string
	ComponentList     string

	list []string
}

// Add ...
func (inst *AutoConfig) Add(comFactoryName string) {
	inst.list = append(inst.list, comFactoryName)
}

////////////////////////////////////////////////////////////////////////////////

// ComplexTemplate ...
type ComplexTemplate struct {
	templateAC *template.Template
}

// MakeAutoConfigFunc 创建 'func autoConfig() {}' 代码片段
func (inst *ComplexTemplate) MakeAutoConfigFunc(ac *AutoConfig) (string, error) {
	t := inst.templateAC
	if t == nil {
		name := "the_auto_config.template"
		raw := theTemplateTextForAutoConfig
		t2, err := template.New(name).Parse(raw)
		if err != nil {
			return "", err
		}
		t = t2
		inst.templateAC = t2
	}
	builder := &strings.Builder{}
	err := t.Execute(builder, ac)
	if err != nil {
		return "", err
	}
	str := builder.String()
	return str, nil
}
