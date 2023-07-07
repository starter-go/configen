package d23gen

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

    
    inst.register(&p0da63b10db_s2_Com2service{})
    inst.register(&pea8e092494_s3_Com3dao{})


    return nil
}
