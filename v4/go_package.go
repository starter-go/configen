package v4

// Package 表示一个 go-package
type Package struct {
	OwnerModule *Module
	FullName    string
	SimpleName  string
	Alias       string
	Path        string
}
