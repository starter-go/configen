package gocode

import "testing"

func TestConfigenParams(t *testing.T) {

	items := []string{}
	items = append(items, "()")
	items = append(items, "(\"value1\")")
	items = append(items, "(a=xyz,b=2,c=3)")
	items = append(items, "(\"#\",\".\")")

	for _, item := range items {
		ct, err := ParseConfigenParams(item)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(ct)
		}
	}
}
