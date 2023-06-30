package readers

import (
	"strings"

	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/gocode"
)

type destinationConfigenGoFileReader struct {
	context *v4.Context
}

func (inst *destinationConfigenGoFileReader) readFiles() error {
	table := inst.context.Destinations
	for _, item := range table {
		err := inst.readFile(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *destinationConfigenGoFileReader) readFile(f *gocode.DestinationFolder) error {
	const (
		filename       = "configen.go"
		prefixPackage1 = "package "
		prefixPackage2 = "package\t"
		prefixConfigen = "//starter:configen"
	)
	file := f.Path.GetChild(filename)
	rows, err := ReadRows(file)
	if err != nil {
		return err
	}

	for _, row := range rows {
		if strings.HasPrefix(row, prefixPackage1) {
			// todo ...
		}
		if strings.HasPrefix(row, prefixPackage2) {
			// todo ...
		}
		if strings.HasPrefix(row, prefixConfigen) {
			// todo ...
		}
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////

// ReadDestinationConfigenGoFiles ...
func ReadDestinationConfigenGoFiles(c *v4.Context) error {
	checker := &destinationConfigenGoFileReader{context: c}
	return checker.readFiles()
}
