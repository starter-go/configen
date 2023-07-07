package d23gen
import (
    pe8a3937f4 "github.com/starter-go/configen/src/example/parts"
    p0da63b10d "github.com/starter-go/configen/src/example/parts/s2"
     "github.com/starter-go/application/components"
)

// type p0da63b10d.Com2service in package:github.com/starter-go/configen/src/example/parts/s2
//
// id:com-0da63b10db169a04-s2-Com2service
// class:class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// alias:alias-e8a3937f481a2a4fcb65cb9f0011b311-IService
// scope:singleton
//
type p0da63b10db_s2_Com2service struct {
}

func (inst* p0da63b10db_s2_Com2service) register(cr components.Registry) {
	r := cr.New()
	r.ID = "com-0da63b10db169a04-s2-Com2service"
	r.Classes = "class-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Aliases = "alias-e8a3937f481a2a4fcb65cb9f0011b311-IService"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	cr.Register(r)
}

func (inst* p0da63b10db_s2_Com2service) new() any {
    return &p0da63b10d.Com2service{}
}

func (inst* p0da63b10db_s2_Com2service) inject(injection components.Injection, instance any) error {
	ie := injection.Ext()
	com := instance.(*p0da63b10d.Com2service)

	
    com.Service = inst.getService(ie)
    com.Controller = inst.getController(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p0da63b10db_s2_Com2service) getService(ie components.InjectionExt)pe8a3937f4.IService{
    return ie.GetComponent("#alias-e8a3937f481a2a4fcb65cb9f0011b311-IService").(pe8a3937f4.IService)
}


func (inst*p0da63b10db_s2_Com2service) getController(ie components.InjectionExt)pe8a3937f4.IController{
    return ie.GetComponent("#alias-e8a3937f481a2a4fcb65cb9f0011b311-IController").(pe8a3937f4.IController)
}


func (inst*p0da63b10db_s2_Com2service) getDao(ie components.InjectionExt)pe8a3937f4.IDao{
    return ie.GetComponent("#alias-e8a3937f481a2a4fcb65cb9f0011b311-IDao").(pe8a3937f4.IDao)
}


