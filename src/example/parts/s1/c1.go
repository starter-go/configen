package s1

import (
	"context"
	"io"

	"github.com/starter-go/configen/src/example/parts"
)

// Com1ctrl ...
type Com1ctrl struct {
	//starter:component
	_as func(parts.IController) //starter:as("#.")

	Service    parts.IService    //starter:inject("#")
	Controller parts.IController //starter:inject("#")
	Dao        parts.IDao        //starter:inject("#")

	REader []io.Reader //starter:inject(".")
}

func (inst *Com1ctrl) _impl() {
	inst._as(inst)
}

// Fetch ...
func (inst *Com1ctrl) Fetch(c context.Context) error {
	return nil
}
