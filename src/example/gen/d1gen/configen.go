package d1gen

import "github.com/starter-go/application/components"

//starter:configen(version="4")

// ConfigForD1  ... 导出自动生成的配置
func ConfigForD1(cr components.Registry) error {
	return registerComponents(cr)
	// return nil
}
