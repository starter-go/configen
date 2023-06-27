package example

import (
	"bytes"
	"crypto/sha1"
	"sort"
	"strings"
)

// Com1 is a demo component for starter configen
//starter:component (id="com1-1",class="com1",scope="singleton",alias="")
type Com1 struct {

	//starter:inject(".")
	Field1 []any

	Field2 []sort.Interface //starter:inject(".")
	Field3 sort.Interface   //starter:inject("#")
	Field4 *strings.Builder //starter:inject("#")
	Field5 *bytes.Buffer    //starter:inject("#")

}

//starter:interface
func (inst *Com1) _Impl() sort.Interface {
	return inst
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
