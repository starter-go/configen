package gen

import "github.com/starter-go/application/components"

//starter:configen(version="4")

// ConfigForConfigenExample ... 导出自动生成的配置
func ConfigForConfigenExample(cr components.Registry) error {
	return autoConfig(cr)
}
