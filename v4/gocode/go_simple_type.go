package gocode

import (
	"fmt"
	"strings"
)

// SimpleType 表示一个简单类型，例如 field 的类型, func-param 的类型
type SimpleType struct {
	Text         string // 完整的表达式
	Words        Words
	Package      Import
	SimpleName   string
	IsPtr        bool // with the '*' mark
	IsNativeType bool // a native type like: int|bool|any|error|string|...
}

// ParseSimpleType 把字符串解析为简单类型
func ParseSimpleType(text string, imports *ImportSet) (*SimpleType, error) {
	words := ParseWords(text)
	return CreateSimpleType(words, imports)
}

// CreateSimpleType 创建简单类型
func CreateSimpleType(words *Words, imports *ImportSet) (*SimpleType, error) {

	st := &SimpleType{}
	st.Words = *words
	st.Text = words.String()
	words2 := words

	// check ptr
	item0 := words.WordAt(0, "")
	if item0 == "*" {
		st.IsPtr = true
		words2 = NewWords(words.list[1:])
	}

	if words2.HasPattern("*") {
		st.SimpleName = words2.WordAt(0, "")
	} else if words2.HasPattern("*", ".", "*") {
		st.Package.Alias = words2.WordAt(0, "")
		st.SimpleName = words2.WordAt(2, "")
	} else {
		return nil, fmt.Errorf("unsupported type: %s", st.Text)
	}

	pkgAlias := st.Package.Alias
	if imports != nil && pkgAlias != "" {
		imp := imports.Find(pkgAlias)
		if imp == nil {
			return nil, fmt.Errorf("no import item with alias: %s", pkgAlias)
		}
		st.Package.FullName = imp.FullName
	}

	// check native type
	simpleName := st.SimpleName
	simpleNameL := strings.ToLower(simpleName)
	st.IsNativeType = (simpleName == simpleNameL)

	return st, nil
}
