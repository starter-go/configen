package gocode

import "github.com/starter-go/afs"

// Source 表示一个 *.go 文件
type Source struct {
	OwnerPackage *Package
	Path         afs.Path //    string
	Name         string

	ImportSet     ImportSet
	TypeStructSet TypeStructSet
}

////////////////////////////////////////////////////////////////////////////////

// SourceList ...
type SourceList struct {
	list []*Source
}

// Add ...
func (inst *SourceList) Add(item *Source) {
	if item == nil || inst == nil {
		return
	}
	inst.list = append(inst.list, item)
}

// List ...
func (inst *SourceList) List() []*Source {
	return inst.list
}
