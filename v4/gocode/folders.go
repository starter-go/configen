package gocode

import (
	"github.com/starter-go/afs"
	"github.com/starter-go/configen/v4/dto"
)

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
