package s1

import "github.com/starter-go/configen/src/example/parts"

// Com1ctrl ...
type Com1ctrl struct {
	//starter:component
	_as func(parts.IController) //starter:as("#.")

	Service    parts.IService    //starter:inject("#")
	Controller parts.IController //starter:inject("#")
	Dao        parts.IDao        //starter:inject("#")
}
