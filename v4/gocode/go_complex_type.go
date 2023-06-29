//xyz

package gocode

import "fmt"

// ComplexType 表示一个复杂类型，例如 map[key]value 的类型, []item 的类型
type ComplexType struct {
	Text      string // 完整的表达式
	Words     Words
	IsArray   bool
	IsMap     bool
	KeyType   *SimpleType
	ValueType SimpleType
}

// ParseComplexType 解析复杂类型
func ParseComplexType(text string, imports *ImportSet) (*ComplexType, error) {
	words := ParseWords(text)
	return CreateComplexType(words, imports)
}

// CreateComplexType 创建复杂类型
func CreateComplexType(words *Words, imports *ImportSet) (*ComplexType, error) {

	ct := &ComplexType{}
	ct.Words = *words
	ct.Text = words.String()

	all := words.list
	part1 := &Words{}
	part2 := &Words{}
	part3 := &Words{}
	ipart := 1
	isCollection := false

	for _, item := range all {
		if item == "[" {
			if ipart == 1 {
				ipart++
			} else {
				return nil, fmt.Errorf("bad complex type: %s", ct.Text)
			}
			isCollection = true
		} else if item == "]" {
			if ipart == 2 {
				ipart++
			} else {
				return nil, fmt.Errorf("bad complex type: %s", ct.Text)
			}
		} else {
			switch ipart {
			case 1:
				part1.list = append(part1.list, item)
				break
			case 2:
				part2.list = append(part2.list, item)
				break
			case 3:
				part3.list = append(part3.list, item)
				break
			default:
				break
			}
		}
	}

	if isCollection {
		p1str := part1.String()
		if p1str == "" {
			// as array
			ct.IsArray = true
		} else if p1str == "map" {
			// as map
			ct.IsMap = true
		} else {
			return nil, fmt.Errorf("bad complex type: %s", ct.Text)
		}
	} else {
		// as simple type
		part2 = part1
		part3 = part1
	}

	st3, err := CreateSimpleType(part3, imports)
	if err != nil {
		return nil, err
	}

	if ct.IsMap {
		st2, err := CreateSimpleType(part2, imports)
		if err != nil {
			return nil, err
		}
		ct.KeyType = st2
	}

	ct.ValueType = *st3
	return ct, nil
}
