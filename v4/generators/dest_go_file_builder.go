package generators

import (
	"crypto/md5"
	"io/fs"
	"os"
	"sort"
	"strings"

	"github.com/starter-go/afs"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/configen/v4/gocode"
	"github.com/starter-go/configen/v4/gocode/codetemplate"
)

////////////////////////////////////////////////////////////////////////////////

type complexGoFileBuilder struct {
	packageSimpleName string
	buffer            strings.Builder
	factoryTypes      []string
	target            afs.Path
}

func (inst *complexGoFileBuilder) makeComponentListText() string {
	//	inst.register(&pe8a3937f48_parts_Com3{})
	builder := strings.Builder{}
	builder.WriteString("\n")
	list := inst.factoryTypes
	sort.Strings(list)
	prevName := ""
	for _, name := range list {
		if prevName != name {
			builder.WriteString("    inst.register(&")
			builder.WriteString(name)
			builder.WriteString("{})\n")
		}
		prevName = name
	}
	return builder.String()
}

func (inst *complexGoFileBuilder) complete() error {

	ac := &codetemplate.AutoConfig{}
	ac.PackageSimpleName = inst.packageSimpleName
	ac.ComponentList = inst.makeComponentListText()

	list := inst.factoryTypes
	for _, name := range list {
		ac.Add(name)
	}

	templ := &codetemplate.ComplexTemplate{}
	text, err := templ.MakeAutoConfigFunc(ac)
	if err != nil {
		return err
	}

	inst.buffer.Reset()
	inst.buffer.WriteString(text)
	return nil
}

func (inst *complexGoFileBuilder) WriteToFile() error {
	err := inst.complete()
	if err != nil {
		return err
	}
	opt := &afs.Options{
		Permission: fs.ModePerm,
		Flag:       os.O_CREATE | os.O_WRONLY,
	}
	text := inst.buffer.String()
	file := inst.target
	return file.GetIO().WriteText(text, opt)
}

func (inst *complexGoFileBuilder) addFactoryType(name string) {
	inst.factoryTypes = append(inst.factoryTypes, name)
}

////////////////////////////////////////////////////////////////////////////////

type simpleGoFileBuilder struct {
	hub               *complexGoFileBuilder
	packageSimpleName string
	buffer            strings.Builder
	template          codetemplate.SimpleTemplate
	target            afs.Path
	completed         bool
	importSet         gocode.ImportSet
}

func (inst *simpleGoFileBuilder) computeInstanceType(ts *gocode.TypeStruct) string {
	pack := ts.OwnerPackage
	imp := inst.ImportPackage(&gocode.Import{
		Alias:    pack.Alias,
		FullName: pack.FullName,
	})
	return imp.Alias + "." + ts.Name
}

func (inst *simpleGoFileBuilder) computeFactoryType(ts *gocode.TypeStruct) string {
	pack := ts.OwnerPackage.FullName
	alias := ts.OwnerPackage.Alias
	name := ts.Name
	sum := md5.Sum([]byte(pack))
	hex := lang.HexFromBytes(sum[0:5])
	return "p" + hex.String() + "_" + alias + "_" + name
}

func (inst *simpleGoFileBuilder) WriteComponent(ts1 *gocode.TypeStruct) error {

	ts2 := &codetemplate.TypeStruct{}
	ts2.ComID = ts1.ComID
	ts2.ComClass = ts1.ComClass
	ts2.ComAlias = ts1.ComAlias
	ts2.ComScope = ts1.ComScope

	ts2.FactoryType = inst.computeFactoryType(ts1)
	ts2.FactoryPackageAlias = inst.packageSimpleName
	ts2.FactoryPackageName = "Todo"

	ts2.InstanceType = inst.computeInstanceType(ts1)
	ts2.InstancePackageAlias = "Todo"
	ts2.InstancePackageName = ts1.OwnerPackage.FullName

	fields := inst.makeFieldInjectorList(ts1, ts2)
	ts2.InjectFieldList = inst.makeInjectFieldList(fields)
	ts2.InjectFieldGetters = inst.makeInjectFieldGetters(fields)

	code, err := inst.template.MakeTypeStruct(ts2)
	if err != nil {
		return err
	}
	b := &inst.buffer
	b.WriteString(code)

	inst.hub.addFactoryType(ts2.FactoryType)
	return nil
}

func (inst *simpleGoFileBuilder) makeFieldInjectorList(ts1 *gocode.TypeStruct, ts2 *codetemplate.TypeStruct) []*FieldInjector {
	factory := ts2.FactoryType
	src := ts1.Fields.List()
	dst := make([]*FieldInjector, 0)
	for _, f1 := range src {

		t1 := &f1.Type
		t2 := inst.ImportComplexType(t1)
		f1.Type = *t2

		f2 := &FieldInjector{}
		f2.init(factory, ts1, f1)
		dst = append(dst, f2)
	}
	return dst
}

func (inst *simpleGoFileBuilder) makeInjectFieldList(fields []*FieldInjector) string {
	builder := strings.Builder{}
	builder.WriteString("\n")
	for _, f := range fields {
		str := f.MakeAssignmentStatement()
		builder.WriteString(str + "\n")
	}
	return builder.String()
}

func (inst *simpleGoFileBuilder) makeInjectFieldGetters(fields []*FieldInjector) string {
	builder := strings.Builder{}
	for _, f := range fields {
		str := f.MakeGetterFunc()
		builder.WriteString(str + "\n")
	}
	return builder.String()
}

func (inst *simpleGoFileBuilder) complete() error {
	if inst.completed {
		return nil
	}

	b := &inst.buffer
	body := b.String()
	b.Reset()
	b.WriteString("package ")
	b.WriteString(inst.packageSimpleName)
	b.WriteString("\n")

	// todo: make import list ...
	inst.makeImportList(b)

	b.WriteString(body)
	inst.completed = true
	return nil
}

func (inst *simpleGoFileBuilder) makeImportList(b *strings.Builder) {
	const (
		mark = "\""
		nl   = "\n"
	)

	src := inst.importSet.List()
	src = inst.addImport(src, "github.com/starter-go/application")

	b.WriteString("import (" + nl)
	for _, item := range src {
		if item.FullName == "" {
			continue
		}
		b.WriteString("    ")
		b.WriteString(item.Alias)
		b.WriteString(" ")
		b.WriteString(mark)
		b.WriteString(item.FullName)
		b.WriteString(mark + nl)
	}
	b.WriteString(")" + nl)
}

func (inst *simpleGoFileBuilder) addImport(list []*gocode.Import, packName string) []*gocode.Import {
	return append(list, &gocode.Import{FullName: packName})
}

func (inst *simpleGoFileBuilder) WriteToFile() error {
	err := inst.complete()
	if err != nil {
		return err
	}
	opt := &afs.Options{
		Permission: fs.ModePerm,
		Flag:       os.O_CREATE | os.O_WRONLY,
	}
	text := inst.buffer.String()
	file := inst.target
	return file.GetIO().WriteText(text, opt)
}

func (inst *simpleGoFileBuilder) ImportComplexType(t *gocode.ComplexType) *gocode.ComplexType {
	if t == nil {
		return nil
	}
	kt := t.KeyType
	vt := &t.ValueType

	if kt != nil {
		kt = inst.ImportSimpleType(kt)
		t.KeyType = kt
	}

	if vt != nil {
		vt = inst.ImportSimpleType(vt)
		t.ValueType = *vt
	}

	return t
}

func (inst *simpleGoFileBuilder) ImportSimpleType(t *gocode.SimpleType) *gocode.SimpleType {
	if t == nil {
		return nil
	}
	p := &t.Package
	p = inst.ImportPackage(p)
	t.Package = *p
	return t
}

func (inst *simpleGoFileBuilder) ImportPackage(t *gocode.Import) *gocode.Import {
	if t == nil {
		return nil
	}
	t.ComputeHexName()
	hex := t.HexName.String()
	t.Alias = "p" + hex[0:9]
	inst.importSet.Add(t)
	return t
}

////////////////////////////////////////////////////////////////////////////////
