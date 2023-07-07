package libconfigen

import (
	"fmt"

	"github.com/starter-go/configen/v4/generators"
)

// Run 运行 configen 工具
func Run() {
	fmt.Println("BitWormhole Starter Configen v4")
	app := &generators.Application{}
	err := app.Run()
	if err != nil {
		panic(err)
	}
}
