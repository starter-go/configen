package s2

import "github.com/starter-go/configen/src/example/parts"

// Com2service ...
type Com2service struct {
	//starter:component
	_as func(parts.IService) //starter:as("#")

	Service    parts.IService    //starter:inject("#")
	Controller parts.IController //starter:inject("#")
	Dao        parts.IDao        //starter:inject("#")
}
