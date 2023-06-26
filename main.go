package main

import (
	"fmt"

	"github.com/starter-go/configen/v4/generators"
	"github.com/starter-go/configen/v4/readers"
	"github.com/starter-go/vlog"
)

func main() {
	vlog.Info("begin")
	fmt.Println("BitWormhole Starter Configen v4")

	app := &generators.Application{}

	app.AddStep(app.LocateWorkingDir)
	app.AddStep(app.LocateGoModule)
	app.AddStep(app.LoadConfigenJSON)
	app.AddStep(readers.ReadGoModuleInfo)
	app.AddStep(generators.LoadSources)
	app.AddStep(generators.LoadDestinations)
	app.AddStep(readers.ReadDestinationConfigenGoFiles)

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
