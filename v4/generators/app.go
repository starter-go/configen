package generators

import (
	v4 "github.com/starter-go/configen/v4"
)

// Application 表示 configen app 本身
type Application struct {
	// steps []func(c *v4.Context) error
}

// Run 应用主入口
func (inst *Application) Run() error {

	// inst.makeSteps()

	slb := &stepListBuilder{}
	steps := slb.makeSteps()
	ctx := &v4.Context{}

	for _, step := range steps {
		if step == nil {
			continue
		}
		err := step(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
