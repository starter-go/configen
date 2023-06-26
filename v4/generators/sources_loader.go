package generators

import (
	"fmt"

	v4 "github.com/starter-go/configen/v4"
	"github.com/starter-go/configen/v4/dto"
)

type sourcesLoader struct {
	context *v4.Context
}

func (inst *sourcesLoader) loadSource(item1 *dto.Source) (*v4.SourceFolder, error) {

	id := item1.ID
	goMod := inst.context.Module.Path
	path := goMod.GetParent().GetChild(item1.Path)

	if !path.IsDirectory() {
		return nil, fmt.Errorf("no configen source directory %s", path.GetPath())
	}

	item2 := &v4.SourceFolder{
		ID:     id,
		Config: *item1,
		Path:   path,
	}
	return item2, nil
}

func (inst *sourcesLoader) load() error {

	ctx := inst.context
	config := ctx.Configuration
	list1 := config.Configen.Sources
	tab2 := make(map[dto.SourceID]*v4.SourceFolder)

	for _, item1 := range list1 {
		id := item1.ID
		item2, err := inst.loadSource(item1)
		if err != nil {
			return err
		}
		tab2[id] = item2
	}

	ctx.Sources = tab2
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// LoadSources 加载来源数据项
func LoadSources(c *v4.Context) error {
	loader := &sourcesLoader{context: c}
	return loader.load()
}
