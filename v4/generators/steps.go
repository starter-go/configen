package generators

import v4 "github.com/starter-go/configen/v4"

type stepFactory struct{}

func (inst *stepFactory) stepToScanSourceFiles() v4.StepFunc {
	step := &goSourceFileScanner{}
	step.init()
	return step.Scan
}

func (inst *stepFactory) stepToCleanDestFolders() v4.StepFunc {
	return nil
}

func (inst *stepFactory) stepToResolveConfigenInfo() v4.StepFunc {
	step := &configenInfoResolve{}
	return step.Resolve
}

func (inst *stepFactory) stepToLogComInfo() v4.StepFunc {
	step := &myComponentInfoLogger{}
	return step.Run
}
