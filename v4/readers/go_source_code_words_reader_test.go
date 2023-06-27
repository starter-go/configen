package readers

import "testing"

func TestGoCodeWordsReader(t *testing.T) {

	rows := []string{
		"package abcdefg",
		"import \"abcdefg\"",
		"type xyz struct {",
		"//starter:component(id=\"example-1\",class=\"example\")",
		"//starter:inject(\"#example-1\")",
		"foo []pppp.Interface   //starter:inject(\"#example-1\")",
		"foo int   //starter:inject(\"#example-1\")",
		"foo pppp.Inter   //  starter : inject (   \"#example-1\") ",
	}

	reader := &GoCodeWordsReader{}
	for _, row := range rows {
		words := reader.Read(row)
		t.Logf("read row: %s", row)
		for i, word := range words {
			t.Logf("  word[%d] = %s", i, word)
		}
	}
}
