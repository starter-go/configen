package generators

import (
	"fmt"
	"strings"

	"github.com/starter-go/configen/v4/gocode"
)

// FieldInjector 具体某个字段的注射器
type FieldInjector struct {
	getterName  string
	fieldName   string
	factoryName string
	fieldType   gocode.ComplexType
	selector    string
}

func (inst *FieldInjector) init(factoryType string, com *gocode.TypeStruct, field *gocode.Field) {
	inst.fieldName = field.Name
	inst.factoryName = factoryType
	inst.getterName = "get" + field.Name
	inst.fieldType = field.Type
	inst.selector = field.Injection
}

// MakeAssignmentStatement 创建注入语句
func (inst *FieldInjector) MakeAssignmentStatement() string {

	builder := strings.Builder{}

	builder.WriteString("    com.")
	builder.WriteString(inst.fieldName)
	builder.WriteString(" = inst.")
	builder.WriteString(inst.getterName)
	builder.WriteString("(ie)")

	return builder.String()
}

// MakeGetterFunc 创建 getter 函数
func (inst *FieldInjector) MakeGetterFunc() string {

	const nl = "\n"
	builder := &strings.Builder{}

	builder.WriteString(nl)
	builder.WriteString("func (inst*")
	builder.WriteString(inst.factoryName)
	builder.WriteString(") ")
	builder.WriteString(inst.getterName)
	builder.WriteString("(ie application.InjectionExt)")
	builder.WriteString(inst.fieldType.String())
	builder.WriteString("{" + nl)
	inst.makeGetterInner(builder)
	builder.WriteString("}" + nl)

	return builder.String()
}

func (inst *FieldInjector) makeGetterInner(b *strings.Builder) {
	t := &inst.fieldType
	if t.IsMap {
		inst.makeGetterInnerForMap(b)
	} else if t.IsArray {
		inst.makeGetterInnerForList(b)
	} else if inst.isContext() {
		inst.makeGetterInnerForContext(b)
	} else {
		inst.makeGetterInnerOne(b)
	}
}

func (inst *FieldInjector) isContext() bool {

	// const (
	// 	wantPackage = "///applictaion"
	// 	wantName    = "Context"
	// )

	// tc := inst.fieldType
	// if tc.IsArray || tc.IsMap {
	// 	return false
	// }
	// ts := tc.ValueType
	// if ts.IsNativeType || ts.IsPtr {
	// 	return false
	// }
	// return (sel == "context") && (ts.Package.FullName == wantPackage) && (ts.SimpleName == wantName)

	sel := inst.selector
	return (sel == "context")
}

func (inst *FieldInjector) makeGetterInnerForContext(b *strings.Builder) {
	// panic("unsupport: inject with a context")
	text := "return ie.GetContext()"
	b.WriteString("    " + text + "\n")
}

func (inst *FieldInjector) makeGetterInnerForMap(b *strings.Builder) {
	panic("unsupport: inject with a map")
}

func (inst *FieldInjector) makeGetterInnerForList(b *strings.Builder) {

	// dst := make([]sort.Interface, 0)
	// src := ie.ListComponents("")
	// for _, item1 := range src {
	// 	item2 := item1.(sort.Interface)
	// 	dst = append(dst, item2)
	// }
	// return dst

	const (
		mkStr = "\""
	)

	sel := mkStr + inst.selector + mkStr
	itemType := inst.fieldType.ValueType.String()
	rows := make([]string, 0)

	rows = append(rows, fmt.Sprintf("dst := make([]%s, 0)", itemType))
	rows = append(rows, fmt.Sprintf("src := ie.ListComponents(%s)", sel))
	rows = append(rows, fmt.Sprintf("for _, item1 := range src {"))
	rows = append(rows, fmt.Sprintf("    item2 := item1.(%s)", itemType))
	rows = append(rows, fmt.Sprintf("    dst = append(dst, item2)"))
	rows = append(rows, fmt.Sprintf("}"))
	rows = append(rows, fmt.Sprintf("return dst"))

	for _, row := range rows {
		b.WriteString("    ")
		b.WriteString(row)
		b.WriteString("\n")
	}
}

func (inst *FieldInjector) makeGetterInnerOne(b *strings.Builder) {
	t := &inst.fieldType
	native := t.ValueType.IsNativeType
	if native {
		inst.makeGetterInnerOneNative(b)
	} else {
		inst.makeGetterInnerOneObject(b)
	}
}

func (inst *FieldInjector) makeGetterInnerOneObject(b *strings.Builder) {

	// b.WriteString("// todo: makeGetterInnerOneObject \n")
	// 	return ie.GetComponent("").(*bytes.Buffer)

	const (
		nl    = "\n"
		mkStr = "\""
	)
	b.WriteString("    ")
	b.WriteString("return ie.GetComponent(" + mkStr)
	b.WriteString(inst.selector)
	b.WriteString(mkStr + ").(")
	b.WriteString(inst.fieldType.String())
	b.WriteString(")" + nl)
}

func (inst *FieldInjector) makeGetterInnerOneNative(b *strings.Builder) {
	// 	return ie.GetInt("")
	const (
		nl    = "\n"
		mkStr = "\""
	)

	tName := inst.fieldType.ValueType.String()
	table := inst.getTypeNameGetterTable()
	method := table[tName]

	if method == "" {
		panic("no getter for native type: " + tName)
	}

	b.WriteString("    return ie.")
	b.WriteString(method)
	b.WriteString("(" + mkStr)
	b.WriteString(inst.selector)
	b.WriteString(mkStr + ")")
	b.WriteString("\n")
}

var theFieldInjectorGetterTable map[string]string

func (inst *FieldInjector) getTypeNameGetterTable() map[string]string {
	table := theFieldInjectorGetterTable
	if table != nil {
		return table
	}
	table = make(map[string]string)

	table["int"] = "GetInt"
	table["int8"] = "GetInt8"
	table["int16"] = "GetInt16"
	table["int32"] = "GetInt32"
	table["int64"] = "GetInt64"

	table["uint"] = "GetUint"
	table["uint8"] = "GetUint8"
	table["uint16"] = "GetUint16"
	table["uint32"] = "GetUint32"
	table["uint64"] = "GetUint64"

	table["bool"] = "GetBool"
	table["float32"] = "GetFloat32"
	table["float64"] = "GetFloat64"
	table["byte"] = "GetByte"
	table["string"] = "GetString"
	table["rune"] = "GetRune"
	table["any"] = "GetAny"

	theFieldInjectorGetterTable = table
	return table
}
