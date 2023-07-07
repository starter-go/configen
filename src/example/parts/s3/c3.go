package s3

import "github.com/starter-go/configen/src/example/parts"

// Com3dao ...
type Com3dao struct {
	//starter:component
	_as func(parts.IDao) //starter:as("#")

	Service    parts.IService      //starter:inject("#")
	Controller []parts.IController //starter:inject(".")
}
