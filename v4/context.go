package v4

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/configen/v4/dto"
	"github.com/starter-go/configen/v4/gocode"
	"github.com/starter-go/configen/v4/vo"
)

// Context ...
type Context struct {
	WD            afs.Path // 工作目录
	Configuration *vo.Configen
	Module        *gocode.Module

	Destinations map[dto.DestinationID]*DestinationFolder
	Sources      map[dto.SourceID]*SourceFolder
}

// StepFunc 定义生成步骤函数
type StepFunc func(c *Context) error

// SourceFolder ...
type SourceFolder struct {
	ID     dto.SourceID
	Config dto.Source
	Path   afs.Path
}

// DestinationFolder ...
type DestinationFolder struct {
	ID      dto.DestinationID
	Config  dto.Destination
	Path    afs.Path
	Sources []*SourceFolder
}
