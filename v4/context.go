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

	Destinations map[dto.DestinationID]*gocode.DestinationFolder
	Sources      map[dto.SourceID]*gocode.SourceFolder
	GoFiles      gocode.SourceList
}

// StepFunc 定义生成步骤函数
type StepFunc func(c *Context) error
