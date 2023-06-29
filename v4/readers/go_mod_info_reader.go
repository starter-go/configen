package readers

import (
	"fmt"
	"strings"

	"github.com/starter-go/afs"
	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/gocode"
)

// ReadGoModuleInfo 读取 go.mod 文件信息
func ReadGoModuleInfo(c *v4.Context) error {

	file := c.Module.Path
	path := file.GetPath()
	fmt.Println("load go module info from ", path)

	reader := &gomodInfoReader{}
	info, err := reader.read(file)
	if err != nil {
		return err
	}

	c.Module = info
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type gomodInfoReader struct {
}

func (inst *gomodInfoReader) read(file afs.Path) (*gocode.Module, error) {
	rows, err := ReadRows(file)
	if err != nil {
		return nil, err
	}
	mod := &gocode.Module{}
	mod.Path = file
	for _, row := range rows {
		res, err := inst.parseModuleRow(row)
		if err != nil {
			return nil, err
		}
		if res[0] == "module" && res[1] != "" {
			mod.Name = res[1]
			return mod, nil
		}
	}
	return nil, fmt.Errorf("bad go.mod file, path=%s", file.GetPath())
}

func (inst *gomodInfoReader) parseModuleRow(row string) ([2]string, error) {

	const (
		ch0 = string(' ')
		ch1 = "\t"
		ch2 = "\n"
	)
	result := [2]string{}
	if row == "" || strings.HasPrefix(row, "//") {
		return result, nil
	}

	row = strings.ReplaceAll(row, ch0, ch2)
	row = strings.ReplaceAll(row, ch1, ch2)
	items := strings.Split(row, ch2)
	i := 0
	for _, item := range items {
		if item == "" {
			continue
		}
		if i == 0 || i == 1 {
			result[i] = item
		} else {
			return result, fmt.Errorf("bad go.mod row: %s", row)
		}
		i++
	}
	return result, nil
}
