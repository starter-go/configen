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
