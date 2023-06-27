package main

import (
	"fmt"

	"github.com/starter-go/configen/v4/generators"
	"github.com/starter-go/vlog"
)

func main() {

	vlog.Info("begin")
	fmt.Println("BitWormhole Starter Configen v4")

	app := &generators.Application{}
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
