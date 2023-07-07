package generators

import (
	"fmt"
	"strings"

	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/gocode"
	"github.com/starter-go/configen/v4/readers"
)

type destDirsChecker struct{}

func (inst *destDirsChecker) run(c *v4.Context) error {
	all := c.Destinations
	for _, dest := range all {
		err := inst.checkDestFolder(dest)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *destDirsChecker) checkDestFolder(f *gocode.DestinationFolder) error {

	file := f.Path.GetChild("configen.go") // 必须有这个文件
	rows, err := readers.ReadRows(file)
	if err != nil {
		return err
	}

	const (
		want         = "//starter:configen(version=\"4\")" // 文件中必须有这一行
		packageSpace = "package "
	)

	packageOk := false
	versionOk := false

	for _, row := range rows {
		row = strings.TrimSpace(row)
		if row == want {
			versionOk = true
		} else if strings.HasPrefix(row, packageSpace) {
			pSimpleName := strings.TrimSpace(row[len(packageSpace):])
			f.PackageSimpleName = pSimpleName
			packageOk = true
		}
		if versionOk && packageOk {
			return nil
		}
	}

	path := file.GetPath()
	return fmt.Errorf("no configen key-row in file [%s], key-row=[%s]", path, want)
}
