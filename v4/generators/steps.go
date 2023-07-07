package generators

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/starter-go/afs/files"
	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/gocode"
	"github.com/starter-go/configen/v4/readers"
	"github.com/starter-go/configen/v4/vo"
)

type stepListBuilder struct {
	steps []func(c *v4.Context) error
}

func (inst *stepListBuilder) addStep(fn v4.StepFunc) {
	if fn != nil {
		list := inst.steps
		list = append(list, fn)
		inst.steps = list
	}
}

func (inst *stepListBuilder) makeSteps() []func(c *v4.Context) error {

	inst.steps = make([]func(c *v4.Context) error, 0)

	// sf := &stepListBuilder{}

	inst.addStep(inst.locateWorkingDir)
	inst.addStep(inst.locateGoModule)
	inst.addStep(inst.loadConfigenJSON)
	inst.addStep(readers.ReadGoModuleInfo)
	inst.addStep(LoadSources)
	inst.addStep(LoadDestinations)
	inst.addStep(readers.ReadDestinationConfigenGoFiles)

	inst.addStep(inst.stepToScanSourceFiles())
	inst.addStep(inst.stepToResolveConfigenInfo())
	inst.addStep(inst.stepToCleanDestFolders())
	inst.addStep(inst.stepToLogComInfo())

	inst.addStep(inst.stepToCheckDestinations())
	inst.addStep(inst.stepToCleanDestinations())
	inst.addStep(inst.stepToMakeDestinations())

	return inst.steps
}

func (inst *stepListBuilder) stepToScanSourceFiles() v4.StepFunc {
	step := &goSourceFileScanner{}
	step.init()
	return step.Scan
}

func (inst *stepListBuilder) stepToCleanDestFolders() v4.StepFunc {
	return nil
}

func (inst *stepListBuilder) stepToResolveConfigenInfo() v4.StepFunc {
	step := &configenInfoResolve{}
	return step.Resolve
}

func (inst *stepListBuilder) stepToLogComInfo() v4.StepFunc {
	step := &myComponentInfoLogger{}
	return step.Run
}

func (inst *stepListBuilder) stepToCheckDestinations() v4.StepFunc {
	step := &destDirsChecker{}
	return step.run
}

func (inst *stepListBuilder) stepToCleanDestinations() v4.StepFunc {
	step := &destDirsCleaner{}
	return step.run
}

func (inst *stepListBuilder) stepToMakeDestinations() v4.StepFunc {
	step := &destFilesMaker{}
	return step.run
}

// LocateWorkingDir ...
func (inst *stepListBuilder) locateWorkingDir(c *v4.Context) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	dir := files.FS().NewPath(wd)
	if !dir.IsDirectory() {
		return fmt.Errorf("the path is not a dir: %s", wd)
	}
	c.WD = dir
	return nil
}

// LocateGoModule ...
func (inst *stepListBuilder) locateGoModule(c *v4.Context) error {
	const forName = "go.mod"
	wd := c.WD
	mod := &gocode.Module{}
	for pdir := wd; pdir != nil; pdir = pdir.GetParent() {
		if pdir.IsDirectory() {
			file := pdir.GetChild(forName)
			if file.IsFile() {
				mod.Path = file
				c.Module = mod
				return nil
			}
		}
	}
	path := wd.GetPath()
	return fmt.Errorf("cannot find 'go.mod' from working dir [%s]", path)
}

// LoadConfigenJSON ...
func (inst *stepListBuilder) loadConfigenJSON(c *v4.Context) error {

	const name = "configen.json"
	file := c.Module.Path.GetParent().GetChild(name)
	data, err := file.GetIO().ReadBinary(nil)
	if err != nil {
		return err
	}

	fmt.Println("load configurations from ", file.GetPath())

	doc := &vo.Configen{}
	err = json.Unmarshal(data, doc)
	if err != nil {
		return err
	}

	// check version
	const versionWant = "4"
	versionHave := doc.Configen.Version
	if versionHave != versionWant {
		return fmt.Errorf("bad configen version, want:%s have:%s", versionWant, versionHave)
	}

	c.Configuration = doc
	return nil
}
