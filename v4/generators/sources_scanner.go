package generators

import (
	"fmt"
	"sort"
	"strings"

	"github.com/starter-go/afs"
	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/gocode"
	"github.com/starter-go/configen/v4/readers"
)

type goSourceFileScanner struct {
	currentContext *v4.Context
	currentSource  *gocode.SourceFolder
	currentDir     afs.Path
	currentFile    afs.Path
}

func (inst *goSourceFileScanner) init() error {
	return nil
}

func (inst *goSourceFileScanner) scanFile(file afs.Path) error {

	name := file.GetName()
	path := file.GetPath()

	if !strings.HasSuffix(name, ".go") {
		// it's not a go source file, skip
		return nil
	}

	inst.currentFile = file
	fmt.Println("scan go source file ", path)

	reader := readers.NewGoSourceFileReader()
	reader.Init(inst.currentContext, inst.currentSource)
	result, err := reader.Read(file)
	if err != nil {
		return err
	}

	// vlog.Debug("", result.Name)
	inst.currentContext.GoFiles.Add(result)
	result.OwnerPackage.OwnerGroup = inst.currentSource
	return nil
}

func (inst *goSourceFileScanner) scanIntoDir(dir afs.Path, r bool, depth int) error {

	const maxDepth = 64
	path := dir.GetPath()

	if depth > maxDepth {
		return fmt.Errorf("the path is too deep, path=%s", path)
	}

	if !dir.IsDirectory() {
		return fmt.Errorf("the path is not a directory, path=%s", path)
	}

	filelist := make([]afs.Path, 0)
	subdirlist := make([]afs.Path, 0)
	namelist := dir.ListNames()
	sort.Strings(namelist)

	for _, name := range namelist {
		child := dir.GetChild(name)
		if child.IsFile() {
			filelist = append(filelist, child)
		} else if child.IsDirectory() {
			subdirlist = append(subdirlist, child)
		}
	}

	for _, file := range filelist {
		err := inst.scanFile(file)
		if err != nil {
			return err
		}
	}

	if r {
		for _, child := range subdirlist {
			err := inst.scanIntoDir(child, r, depth+1)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (inst *goSourceFileScanner) scanSourceDir(folder *gocode.SourceFolder) error {
	inst.currentSource = folder
	dir := folder.Path
	r := folder.Config.Recursive
	fmt.Printf("scan source tree:%s (at %s)\n", folder.ID, dir.GetPath())
	return inst.scanIntoDir(dir, r, 0)
}

func (inst *goSourceFileScanner) Scan(c *v4.Context) error {
	inst.currentContext = c
	srcList := c.Sources
	for _, srcDir := range srcList {
		err := inst.scanSourceDir(srcDir)
		if err != nil {
			return err
		}
	}
	return nil
}
