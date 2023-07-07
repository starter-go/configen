package s2

import (
	"context"

	"github.com/starter-go/application"
	"github.com/starter-go/configen/src/example/parts"
	"github.com/starter-go/vlog"
)

// Com2service ...
type Com2service struct {
	//starter:component
	_as func(parts.IService, application.Lifecycle) //starter:as("#",".")

	Service    parts.IService    //starter:inject("#")
	Controller parts.IController //starter:inject("#")
	Dao        parts.IDao        //starter:inject("#")
}

func (inst *Com2service) _impl() {
	inst._as(inst, inst)
}

// Fetch ...
func (inst *Com2service) Fetch(c context.Context, id string) (string, error) {
	return "", nil
}

// Life ...
func (inst *Com2service) Life() *application.Life {
	return &application.Life{
		OnCreate:  inst.onCreate,
		OnStart:   inst.onStart,
		OnStop:    inst.onStop,
		OnDestroy: inst.onDestroy,
	}
}

func (inst *Com2service) onCreate() error {
	vlog.Warn("%v.onCreate", inst)
	return nil
}

func (inst *Com2service) onStart() error {
	vlog.Warn("%v.onStart", inst)
	return nil
}

func (inst *Com2service) onStop() error {
	vlog.Warn("%v.onStop", inst)
	return nil
}

func (inst *Com2service) onDestroy() error {
	vlog.Warn("%v.onDestroy", inst)
	return nil
}
