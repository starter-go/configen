package parts

import (
	"io"
)

// Com2 ...
type Com2 struct {

	//starter:component
	_as func(io.Writer, io.Reader) //starter:as(".","#.")

	F11 int   //starter:inject("${a.b.c.d}")
	F12 int8  //starter:inject("${a.b.c.d}")
	F13 int16 //starter:inject("${a.b.c.d}")
	F14 int32 //starter:inject("${a.b.c.d}")
	F15 int64 //starter:inject("${a.b.c.d}")

	F21 uint   //starter:inject("${a.b.c.d}")
	F22 uint8  //starter:inject("${a.b.c.d}")
	F23 uint16 //starter:inject("${a.b.c.d}")
	F24 uint32 //starter:inject("${a.b.c.d}")
	F25 uint64 //starter:inject("${a.b.c.d}")

	F31 string  //starter:inject("${a.b.c.d}")
	F32 byte    //starter:inject("${a.b.c.d}")
	F33 bool    //starter:inject("${a.b.c.d}")
	F34 rune    //starter:inject("${a.b.c.d}")
	F35 float32 //starter:inject("${a.b.c.d}")
	F36 float64 //starter:inject("${a.b.c.d}")
	F37 any     //starter:inject("${a.b.c.d}")
}

func (inst *Com2) _impl() {
	inst._as(inst, inst)
}

func (inst *Com2) Read(data []byte) (int, error) {
	return 0, io.EOF
}

func (inst *Com2) Write(data []byte) (int, error) {
	return 0, io.EOF
}
