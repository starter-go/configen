package s3

import (
	"context"

	"github.com/starter-go/application"
	"github.com/starter-go/configen/src/example/parts"
)

////////////////////////////////////////////////////////////////////////////////

// Com3dao ...
type Com3dao struct {
	//starter:component
	_as func(parts.IDao) //starter:as("#")

	Service    parts.IService      //starter:inject("#")
	Controller []parts.IController //starter:inject(".")
}

func (inst *Com3dao) _impl() {
	inst._as(inst)
}

// Fetch ...
func (inst *Com3dao) Fetch(c context.Context, id int) (string, error) {
	return "", nil
}

////////////////////////////////////////////////////////////////////////////////

// Com3x  ...
type Com3x struct {
	//starter:component
	_as func(parts.IDao) //starter:as(".")

	Context application.Context //starter:inject("context")

}

////////////////////////////////////////////////////////////////////////////////

// Com3z  ...
type Com3z struct {
	//starter:component
	_as func(parts.IDao) //starter:as(".")

	// Service    parts.IService      //starter:inject("#")
	// Controller []parts.IController //starter:inject(".")
}

////////////////////////////////////////////////////////////////////////////////
