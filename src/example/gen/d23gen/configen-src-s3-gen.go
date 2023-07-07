package d23gen
import (
    pea8e09249 "github.com/starter-go/configen/src/example/parts/s3"
    pe8a3937f4 "github.com/starter-go/configen/src/example/parts"
     "github.com/starter-go/application/components"
)

// type pea8e09249.Com3dao in package:github.com/starter-go/configen/src/example/parts/s3
//
// id:com-ea8e092494300fad-s3-Com3dao
// class:
// alias:alias-e8a3937f481a2a4fcb65cb9f0011b311-IDao
// scope:singleton
//
type pea8e092494_s3_Com3dao struct {
}

func (inst* pea8e092494_s3_Com3dao) register(cr components.Registry) {
	r := cr.New()
	r.ID = "com-ea8e092494300fad-s3-Com3dao"
	r.Classes = ""
	r.Aliases = "alias-e8a3937f481a2a4fcb65cb9f0011b311-IDao"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	cr.Register(r)
}

func (inst* pea8e092494_s3_Com3dao) new() any {
    return &pea8e09249.Com3dao{}
}

func (inst* pea8e092494_s3_Com3dao) inject(injection components.Injection, instance any) error {
	ie := injection.Ext()
	com := instance.(*pea8e09249.Com3dao)

	
    com.Service = inst.getService(ie)
    com.Controller = inst.getController(ie)


    return nil
}


func (inst*pea8e092494_s3_Com3dao) getService(ie components.InjectionExt)pe8a3937f4.IService{
    return ie.GetComponent("#alias-e8a3937f481a2a4fcb65cb9f0011b311-IService").(pe8a3937f4.IService)
}


func (inst*pea8e092494_s3_Com3dao) getController(ie components.InjectionExt)[]pe8a3937f4.IController{
    dst := make([]pe8a3937f4.IController, 0)
    src := ie.ListComponents(".class-e8a3937f481a2a4fcb65cb9f0011b311-IController")
    for _, item1 := range src {
        item2 := item1.(pe8a3937f4.IController)
        dst = append(dst, item2)
    }
    return dst
}


