package d1gen
import (
    p2fc3c2a45 "github.com/starter-go/configen/src/example/parts/s1"
    pe8a3937f4 "github.com/starter-go/configen/src/example/parts"
     "github.com/starter-go/application/components"
)

// type p2fc3c2a45.Com1ctrl in package:github.com/starter-go/configen/src/example/parts/s1
//
// id:com-2fc3c2a45d4e6b8e-s1-Com1ctrl
// class:com-e8a3937f481a2a4fcb65cb9f0011b311-IController
// alias:com-e8a3937f481a2a4fcb65cb9f0011b311-IController
// scope:singleton
//
type p2fc3c2a45d_s1_Com1ctrl struct {
}

func (inst* p2fc3c2a45d_s1_Com1ctrl) register(cr components.Registry) {
	r := cr.New()
	r.ID = "com-2fc3c2a45d4e6b8e-s1-Com1ctrl"
	r.Classes = "com-e8a3937f481a2a4fcb65cb9f0011b311-IController"
	r.Aliases = "com-e8a3937f481a2a4fcb65cb9f0011b311-IController"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	cr.Register(r)
}

func (inst* p2fc3c2a45d_s1_Com1ctrl) new() any {
    return &p2fc3c2a45.Com1ctrl{}
}

func (inst* p2fc3c2a45d_s1_Com1ctrl) inject(injection components.Injection, instance any) error {
	ie := injection.Ext()
	com := instance.(*p2fc3c2a45.Com1ctrl)

	    com.Service = inst.getService(ie)
    com.Controller = inst.getController(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p2fc3c2a45d_s1_Com1ctrl) getService(ie components.InjectionExt)pe8a3937f4.IService{
return ie.GetComponent("#id-e8a3937f481a2a4fcb65cb9f0011b311-IService").(pe8a3937f4.IService)}


func (inst*p2fc3c2a45d_s1_Com1ctrl) getController(ie components.InjectionExt)pe8a3937f4.IController{
return ie.GetComponent("#id-e8a3937f481a2a4fcb65cb9f0011b311-IController").(pe8a3937f4.IController)}


func (inst*p2fc3c2a45d_s1_Com1ctrl) getDao(ie components.InjectionExt)pe8a3937f4.IDao{
return ie.GetComponent("#id-e8a3937f481a2a4fcb65cb9f0011b311-IDao").(pe8a3937f4.IDao)}


