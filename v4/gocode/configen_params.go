package gocode

import "fmt"

// ConfigenParam 表示一项配置参数
type ConfigenParam struct {
	Name  string
	Value string
}

// ConfigenParams 表示嵌入在 go 代码注释中的配置参数
type ConfigenParams struct {
	items []*ConfigenParam
}

// Add ...
func (inst *ConfigenParams) Add(item *ConfigenParam) *ConfigenParams {
	if item != nil {
		inst.items = append(inst.items, item)
	}
	return inst
}

// GetItems ...
func (inst *ConfigenParams) GetItems() []*ConfigenParam {
	list := inst.items
	if list == nil {
		list = make([]*ConfigenParam, 0)
	}
	return list
}

////////////////////////////////////////////////////////////////////////////////

// ParseConfigenParams 解析配置参数
func ParseConfigenParams(s string) (*ConfigenParams, error) {
	w := ParseWords(s)
	return CreateConfigenParams(w)
}

// CreateConfigenParams 创建配置参数
func CreateConfigenParams(w *Words) (*ConfigenParams, error) {

	elements := w.List()
	lastIndex := len(elements) - 1
	params := &ConfigenParams{}
	p := &ConfigenParam{}

	for i := 0; i <= lastIndex; i++ {
		el := elements[i]

		if i == 0 {
			if el == "(" {
				continue
			} else {
				return nil, fmt.Errorf("bad configen params: %s", w.String())
			}
		} else if i == lastIndex {
			if el == ")" {
				params.Add(p)
				break
			} else {
				return nil, fmt.Errorf("bad configen params: %s", w.String())
			}
		}

		if el == "," {
			params.Add(p)
			p = &ConfigenParam{}
		} else if el == "=" {
			p.Name = p.Value
			p.Value = ""
		} else {
			p.Value = el
		}
	}

	return params, nil
}
