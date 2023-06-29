package gocode

import "github.com/starter-go/afs"

// Package 表示一个 go-package
type Package struct {
	OwnerModule *Module
	FullName    string
	SimpleName  string
	Alias       string
	Path        afs.Path // string
}
