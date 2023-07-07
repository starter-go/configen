package d1gen
import (
    p2fc3c2a45 "github.com/starter-go/configen/src/example/parts/s1"
    pe8a3937f4 "github.com/starter-go/configen/src/example/parts"
    pf98ed07a4 "io"
     "github.com/starter-go/application"
)

// type p2fc3c2a45.Com1ctrl in package:github.com/starter-go/configen/src/example/parts/s1
//
// id:com-2fc3c2a45d4e6b8e-s1-Com1ctrl
// class:class-e8a3937f481a2a4fcb65cb9f0011b311-IController
// alias:alias-e8a3937f481a2a4fcb65cb9f0011b311-IController
// scope:singleton
//
type p2fc3c2a45d_s1_Com1ctrl struct {
}

func (inst* p2fc3c2a45d_s1_Com1ctrl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2fc3c2a45d4e6b8e-s1-Com1ctrl"
	r.Classes = "class-e8a3937f481a2a4fcb65cb9f0011b311-IController"
	r.Aliases = "alias-e8a3937f481a2a4fcb65cb9f0011b311-IController"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2fc3c2a45d_s1_Com1ctrl) new() any {
    return &p2fc3c2a45.Com1ctrl{}
}

func (inst* p2fc3c2a45d_s1_Com1ctrl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2fc3c2a45.Com1ctrl)
	nop(ie, com)

	
    com.Service = inst.getService(ie)
    com.Controller = inst.getController(ie)
    com.Dao = inst.getDao(ie)
    com.REader = inst.getREader(ie)


    return nil
}


func (inst*p2fc3c2a45d_s1_Com1ctrl) getService(ie application.InjectionExt)pe8a3937f4.IService{
    return ie.GetComponent("#alias-e8a3937f481a2a4fcb65cb9f0011b311-IService").(pe8a3937f4.IService)
}


func (inst*p2fc3c2a45d_s1_Com1ctrl) getController(ie application.InjectionExt)pe8a3937f4.IController{
    return ie.GetComponent("#alias-e8a3937f481a2a4fcb65cb9f0011b311-IController").(pe8a3937f4.IController)
}


func (inst*p2fc3c2a45d_s1_Com1ctrl) getDao(ie application.InjectionExt)pe8a3937f4.IDao{
    return ie.GetComponent("#alias-e8a3937f481a2a4fcb65cb9f0011b311-IDao").(pe8a3937f4.IDao)
}


func (inst*p2fc3c2a45d_s1_Com1ctrl) getREader(ie application.InjectionExt)[]pf98ed07a4.Reader{
    dst := make([]pf98ed07a4.Reader, 0)
    src := ie.ListComponents(".class-f98ed07a4d5f50f7de1410d905f1477f-Reader")
    for _, item1 := range src {
        item2 := item1.(pf98ed07a4.Reader)
        dst = append(dst, item2)
    }
    return dst
}


