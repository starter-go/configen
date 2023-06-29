package generators

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/starter-go/afs/files"
	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/gocode"
	"github.com/starter-go/configen/v4/vo"
)

// LocateWorkingDir ...
func (inst *Application) locateWorkingDir(c *v4.Context) error {
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
func (inst *Application) locateGoModule(c *v4.Context) error {
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
func (inst *Application) loadConfigenJSON(c *v4.Context) error {

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
