package d1gen

import "github.com/starter-go/application/components"

func registerComponents(cr components.Registry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr components.Registry)
}

type autoRegistrar struct {
    cr components.Registry
}

func (inst *autoRegistrar) init(cr components.Registry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) {
	factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&p2fc3c2a45d_s1_Com1ctrl{})


    return nil
}
