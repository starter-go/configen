package generators

import (
	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/readers"
)

// Application 表示 configen app 本身
type Application struct {
	steps []func(c *v4.Context) error
}

// Run 应用主入口
func (inst *Application) Run() error {

	inst.makeSteps()
	steps := inst.steps
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

func (inst *Application) addStep(fn v4.StepFunc) {
	if fn != nil {
		list := inst.steps
		list = append(list, fn)
		inst.steps = list
	}
}

func (inst *Application) makeSteps() {

	sf := &stepFactory{}

	inst.addStep(inst.locateWorkingDir)
	inst.addStep(inst.locateGoModule)
	inst.addStep(inst.loadConfigenJSON)
	inst.addStep(readers.ReadGoModuleInfo)
	inst.addStep(LoadSources)
	inst.addStep(LoadDestinations)
	inst.addStep(readers.ReadDestinationConfigenGoFiles)

	inst.addStep(sf.stepToScanSourceFiles())
	inst.addStep(sf.stepToResolveConfigenInfo())
	inst.addStep(sf.stepToCleanDestFolders())
	inst.addStep(sf.stepToLogComInfo())

}
