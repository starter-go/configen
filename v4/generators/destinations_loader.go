package generators

import (
	"fmt"

	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/dto"
	"github.com/starter-go/configen/v4/gocode"
)

type destinationsLoader struct {
	context *v4.Context
}

func (inst *destinationsLoader) loadDestination(item1 *dto.Destination) (*gocode.DestinationFolder, error) {

	id := item1.ID
	goMod := inst.context.Module.Path
	path := goMod.GetParent().GetChild(item1.Path)

	if !path.IsDirectory() {
		return nil, fmt.Errorf("no configen destination directory %s", path.GetPath())
	}

	item2 := &gocode.DestinationFolder{
		ID:     id,
		Config: *item1,
		Path:   path,
	}

	//for sources
	listSrcIDs := item1.Sources
	srcTable := inst.context.Sources
	for _, sid := range listSrcIDs {
		src := srcTable[sid]
		if src == nil {
			return nil, fmt.Errorf("no configen source with id: %s", sid)
		}
		item2.Sources = append(item2.Sources, src)
	}

	return item2, nil
}

func (inst *destinationsLoader) load() error {

	ctx := inst.context
	config := ctx.Configuration
	list1 := config.Configen.Destinations
	tab2 := make(map[dto.DestinationID]*gocode.DestinationFolder)

	for _, item1 := range list1 {
		id := item1.ID
		item2, err := inst.loadDestination(item1)
		if err != nil {
			return err
		}
		tab2[id] = item2
	}

	ctx.Destinations = tab2
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// LoadDestinations 加载目标位置数据项
func LoadDestinations(c *v4.Context) error {
	loader := &destinationsLoader{context: c}
	return loader.load()
}
