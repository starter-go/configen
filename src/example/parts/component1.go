package parts

import (
	"bytes"
	"context"
	"crypto/sha1"
	"io"
	"sort"
	"strings"

	"github.com/starter-go/application"
)

// Com1 is a demo component for starter configen
type Com1 struct {

	//starter:component (id="com1-1",class="com1",scope="singleton",alias="")
	_as func(sort.Interface, context.Context) //starter:as ( "#.", ".")

	Field1 []any //starter:inject(".")

	Field2 []sort.Interface    //starter:inject(".")
	Field3 sort.Interface      //starter:inject("#")
	Field4 *strings.Builder    //starter:inject("#")
	Field5 *bytes.Buffer       //starter:inject("#")
	Field6 io.Reader           //starter:inject("#")
	Field7 application.Context //starter:inject("context")
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
