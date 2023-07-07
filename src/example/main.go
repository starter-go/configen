package main

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/application/boot"
	"github.com/starter-go/configen/src/example/gen"
)

func main() {
	m := theModule()
	opt := &boot.Options{}
	err := boot.Run(m, opt)
	if err != nil {
		panic(err)
	}
}

const (
	theModuleName     = "configen/example1"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
	theModuleResPath  = "res"
)

//go:embed "res"
var theModuleResFS embed.FS

func theModule() application.Module {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName).Version(theModuleVersion).Revision(theModuleRevision)
	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen.ConfigForConfigenExample)
	return mb.Create()
}
