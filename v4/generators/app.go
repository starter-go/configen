package generators

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/starter-go/afs/files"
	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/vo"
)

// Application 表示 configen app 本身
type Application struct {
	steps []func(c *v4.Context) error
}

// AddStep ...
func (inst *Application) AddStep(fn func(c *v4.Context) error) *Application {
	if fn != nil {
		inst.steps = append(inst.steps, fn)
	}
	return inst
}

// Run 应用主入口
func (inst *Application) Run() error {

	ctx := &v4.Context{}
	steps := inst.steps

	for _, step := range steps {
		err := step(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// LocateWorkingDir ...
func (inst *Application) LocateWorkingDir(c *v4.Context) error {
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
func (inst *Application) LocateGoModule(c *v4.Context) error {
	const forName = "go.mod"
	wd := c.WD
	mod := &v4.Module{}
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
func (inst *Application) LoadConfigenJSON(c *v4.Context) error {

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
