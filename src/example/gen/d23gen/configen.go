package d23gen

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComponents  ... 导出自动生成的配置
func ExportComponents(cr application.ComponentRegistry) error {
	return registerComponents(cr)
}
