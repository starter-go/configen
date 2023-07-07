package gen

import "github.com/starter-go/application/components"

func autoConfig(cr components.Registry) error {
    ac:=&autoConfiger{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr components.Registry)
}

type autoConfiger struct {
    cr components.Registry
}

func (inst *autoConfiger) init(cr components.Registry) {
	inst.cr = cr
}

func (inst *autoConfiger) register(factory comFactory) {
	factory.register(inst.cr)
}

func (inst*autoConfiger) addAll() error {

    
    inst.register(&pe8a3937f48_parts_Com1{})
    inst.register(&pe8a3937f48_parts_Com2{})
    inst.register(&pe8a3937f48_parts_Com3{})


    return nil
}
