package main

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/application/boot"
	"github.com/starter-go/configen/src/example/gen/d1gen"
	"github.com/starter-go/configen/src/example/gen/d23gen"
)

func main() {
	m := theModule1()
	opt := &boot.Options{}
	err := boot.Run(m, opt)
	if err != nil {
		panic(err)
	}
}

const (
	theModuleName     = "github.com/starter-go/configen/example"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
	theModuleResPath  = "res"
)

//go:embed "res"
var theModuleResFS embed.FS

func theModule1() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "-m1").Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(d1gen.ConfigForD1)
	mb.Depend(theModule23())
	return mb.Create()
}

func theModule23() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "-m23").Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(d23gen.ExportComponents)
	return mb.Create()
}
