package generators

import (
	"strings"

	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/gocode"
)

type configenInfoResolve struct {
}

func (inst *configenInfoResolve) Resolve(c *v4.Context) error {
	list := c.GoFiles.List()
	for _, gofile := range list {
		err := inst.resolveGoFile(c, gofile)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *configenInfoResolve) resolveGoFile(c *v4.Context, file *gocode.Source) error {
	list := file.TypeStructSet.List()
	for _, ts := range list {
		err := inst.resolveTypeStruct(c, ts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *configenInfoResolve) resolveTypeStruct(c *v4.Context, ts *gocode.TypeStruct) error {

	// field-list
	list := ts.Fields.List()
	for _, field := range list {
		err := inst.resolveField(c, field)
		if err != nil {
			return err
		}
	}

	// as-list
	list2 := ts.As.List()
	for _, impl := range list2 {
		err := inst.resolveComImpl(c, ts, impl)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *configenInfoResolve) resolveComImpl(c *v4.Context, ts *gocode.TypeStruct, impl *gocode.Implementation) error {

	selector := impl.Injection

	if strings.ContainsRune(selector, '#') {
		// add to alias
		inj := inst.makeInjectionNameForType(&impl.Type, '#')
		ts.ComAlias = strings.TrimSpace(ts.ComAlias + " " + inj)
	}

	if strings.ContainsRune(selector, '.') {
		// add to class
		inj := inst.makeInjectionNameForType(&impl.Type, '.')
		ts.ComClass = strings.TrimSpace(ts.ComClass + " " + inj)
	}

	return nil
}

func (inst *configenInfoResolve) resolveField(c *v4.Context, f *gocode.Field) error {
	selector := f.Injection
	if selector == "#" || selector == "." {
		ct := &f.Type
		selectorRuneList := []rune(selector)
		inj := inst.makeInjectionNameForType(ct, selectorRuneList[0])
		f.Injection = selector + inj
	}
	return nil
}

// makeInjectionNameForType: rawSel=['#'|'.']
func (inst *configenInfoResolve) makeInjectionNameForType(ct *gocode.ComplexType, rawSel rune) string {
	vt := ct.ValueType
	pkg := vt.Package
	fullname := pkg.FullName
	hex := pkg.HexName
	if hex == "" {
		pkg.ComputeHexName()
		hex = pkg.HexName
	}
	if fullname == "" {
		if vt.IsNativeType {
			hex = "native"
		} else {
			panic("the package full-name is empty")
		}
	}
	prefix := "com-"
	if rawSel == '#' {
		prefix = "alias-"
	} else if rawSel == '.' {
		prefix = "class-"
	}
	return prefix + hex.String() + "-" + vt.SimpleName
}
