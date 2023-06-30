package generators

import (
	"fmt"
	"strings"

	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/gocode"
)

type myComponentInfoLogger struct{}

func (inst *myComponentInfoLogger) Run(c *v4.Context) error {
	srcfiles := c.GoFiles.List()
	for _, srcfile := range srcfiles {
		inst.logGoFile(srcfile)
	}
	return nil
}

func (inst *myComponentInfoLogger) logGoFile(file *gocode.Source) {
	comlist := file.TypeStructSet.List()
	for _, com := range comlist {
		inst.logCom(com)
	}
}

func (inst *myComponentInfoLogger) logCom(com *gocode.TypeStruct) {
	const tab = "\n\t"
	builder := &strings.Builder{}

	builder.WriteString("component:")
	builder.WriteString(com.ComID)

	// type
	builder.WriteString(tab + "type:")
	builder.WriteString(com.Name)
	builder.WriteString(" (@")
	builder.WriteString(com.OwnerPackage.FullName)
	builder.WriteString(")")

	// id
	builder.WriteString(tab + "id:" + com.ComID)

	// class
	builder.WriteString(tab + "class:" + com.ComClass)

	// alias
	builder.WriteString(tab + "alias:" + com.ComAlias)

	// scope
	builder.WriteString(tab + "scope:" + com.ComScope)

	builder.WriteString(tab + "as:")
	inst.logComAsList(com, builder)

	builder.WriteString(tab + "injections:")
	inst.logComFields(com, builder)

	fmt.Println(builder.String())
}

func (inst *myComponentInfoLogger) logComAsList(com *gocode.TypeStruct, b *strings.Builder) {

	list := com.As.List()
	for _, impl := range list {
		b.WriteString("\n\t\t")

		b.WriteString(impl.Name)
		b.WriteString("  ")
		b.WriteString(impl.Type.Words.String())

		b.WriteString(" //as:'")
		b.WriteString(impl.Injection)
		b.WriteString("'")

	}
}

func (inst *myComponentInfoLogger) logComFields(com *gocode.TypeStruct, b *strings.Builder) {
	list := com.Fields.List()
	for _, field := range list {
		b.WriteString("\n\t\t")
		b.WriteString(field.Name)
		b.WriteString("  ")
		b.WriteString(field.Type.Words.String())
		b.WriteString(" //inject:'")
		b.WriteString(field.Injection)
		b.WriteString("'")
	}
}
