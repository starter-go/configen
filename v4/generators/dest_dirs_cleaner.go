package generators

import (
	"strings"

	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/gocode"
	"github.com/starter-go/vlog"
)

type destDirsCleaner struct{}

func (inst *destDirsCleaner) run(c *v4.Context) error {
	all := c.Destinations
	for _, dest := range all {
		err := inst.cleanDir(dest)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *destDirsCleaner) cleanDir(f *gocode.DestinationFolder) error {

	const (
		prefix = "configen-"
		suffix = "-gen.go"
	)

	dir := f.Path
	path := dir.GetPath()
	list := dir.ListChildren()

	vlog.Info("clean destination folder: %s", path)

	for _, child := range list {
		name := child.GetName()
		if strings.HasPrefix(name, prefix) && strings.HasSuffix(name, suffix) {
			if child.IsFile() {
				vlog.Info("  delete file %s", name)
				child.Delete()
			}
		}
	}

	return nil
}
