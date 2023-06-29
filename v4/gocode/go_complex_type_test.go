package gocode

import "testing"

func TestComplexType(t *testing.T) {

	items := []string{}
	items = append(items, "int")
	items = append(items, "sort.Interface")
	items = append(items, "*strings.Builder")
	items = append(items, "[]*strings.Builder")
	items = append(items, "map[string]*strings.Builder")

	for _, item := range items {
		ct, err := ParseComplexType(item, nil)
		if err != nil {
			t.Error(err)
		} else {
			t.Log(ct)
		}
	}
}
