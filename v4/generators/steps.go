package generators

import v4 "github.com/starter-go/configen/v4"

type stepFactory struct{}

func (inst *stepFactory) stepToScanSourceFiles() v4.StepFunc {
	step := &goSourceFileScanner{}
	step.init()
	return step.Scan
}
