package v4

import "github.com/starter-go/afs"

// Module 表示一个 go-module
type Module struct {
	Name string   // the full-name of module
	Path afs.Path // path of 'go.mod'
}
