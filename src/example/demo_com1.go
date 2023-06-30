package example

import (
	"bytes"
	"context"
	"crypto/sha1"
	"io"
	"sort"
	"strings"
)

// Com1 is a demo component for starter configen
type Com1 struct {

	//starter:component (id="com1-1",class="com1",scope="singleton",alias="")
	_as func(sort.Interface, context.Context) //starter:as ( "#.", ".")

	Field1 []any //starter:inject(".")

	Field2 []sort.Interface //starter:inject(".")
	Field3 sort.Interface   //starter:inject("#")
	Field4 *strings.Builder //starter:inject("#")
	Field5 *bytes.Buffer    //starter:inject("#")
	Field6 io.Reader        //starter:inject("#")
}

func (inst *Com1) _impl() {
	inst._as(inst, nil)
}

func (inst *Com1) Len() int {

	data := ""
	sha1.Sum([]byte(data))

	return 0
}

func (inst *Com1) Less(a, b int) bool {
	return false
}

func (inst *Com1) Swap(a, b int) {
}

////////////////////////////////////////////////////////////////////////////////

// Com2 ...
type Com2 struct {

	//starter:component
	_as func(io.Writer, io.Reader) //starter:as(".","#.")

	F1 int     //starter:inject("${a.b.c.d}")
	F2 *ComNot //starter:inject("#com-not")

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

////////////////////////////////////////////////////////////////////////////////

// Com3  ...
type Com3 struct {

	//starter:component("com-3-abc", alias="c3-xyz c3-ijk")

	A int //starter:inject("#")
	B bool
	C rune
	D []byte
	E string
	F float32
}

////////////////////////////////////////////////////////////////////////////////

// ComNot 这不是一个组件
type ComNot struct {
	A int //starter:inject("666")
	B bool
	C rune
	D []byte
	E string
	F float32
}
