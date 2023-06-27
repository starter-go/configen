package v4

import "github.com/starter-go/afs"

// Source 表示一个 go-source-file
type Source struct {
	OwnerPackage *Package
	Path         afs.Path //    string
	Name         string
}

type CodeImport struct {
	Alias    string
	FullName string
}

type CodeTypeStruct struct {
	Name string
}

type CodeStructField struct {
	Name string
}

type CodeStructFunc struct {
	Name string
}
