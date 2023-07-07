package d23gen

import "github.com/starter-go/application/components"

//starter:configen(version="4")

// ExportComponents  ... 导出自动生成的配置
func ExportComponents(cr components.Registry) error {
	return registerComponents(cr)
}
