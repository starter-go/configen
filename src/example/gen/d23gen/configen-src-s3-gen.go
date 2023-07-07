package d23gen
import (
    pea8e09249 "github.com/starter-go/configen/src/example/parts/s3"
    pe8a3937f4 "github.com/starter-go/configen/src/example/parts"
    p0ef6f2938 "github.com/starter-go/application"
     "github.com/starter-go/application"
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

func (inst* pea8e092494_s3_Com3dao) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ea8e092494300fad-s3-Com3dao"
	r.Classes = ""
	r.Aliases = "alias-e8a3937f481a2a4fcb65cb9f0011b311-IDao"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pea8e092494_s3_Com3dao) new() any {
    return &pea8e09249.Com3dao{}
}

func (inst* pea8e092494_s3_Com3dao) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pea8e09249.Com3dao)
	nop(ie, com)

	
    com.Service = inst.getService(ie)
    com.Controller = inst.getController(ie)


    return nil
}


func (inst*pea8e092494_s3_Com3dao) getService(ie application.InjectionExt)pe8a3937f4.IService{
    return ie.GetComponent("#alias-e8a3937f481a2a4fcb65cb9f0011b311-IService").(pe8a3937f4.IService)
}


func (inst*pea8e092494_s3_Com3dao) getController(ie application.InjectionExt)[]pe8a3937f4.IController{
    dst := make([]pe8a3937f4.IController, 0)
    src := ie.ListComponents(".class-e8a3937f481a2a4fcb65cb9f0011b311-IController")
    for _, item1 := range src {
        item2 := item1.(pe8a3937f4.IController)
        dst = append(dst, item2)
    }
    return dst
}



// type pea8e09249.Com3x in package:github.com/starter-go/configen/src/example/parts/s3
//
// id:com-ea8e092494300fad-s3-Com3x
// class:class-e8a3937f481a2a4fcb65cb9f0011b311-IDao
// alias:
// scope:singleton
//
type pea8e092494_s3_Com3x struct {
}

func (inst* pea8e092494_s3_Com3x) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ea8e092494300fad-s3-Com3x"
	r.Classes = "class-e8a3937f481a2a4fcb65cb9f0011b311-IDao"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pea8e092494_s3_Com3x) new() any {
    return &pea8e09249.Com3x{}
}

func (inst* pea8e092494_s3_Com3x) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pea8e09249.Com3x)
	nop(ie, com)

	
    com.Context = inst.getContext(ie)


    return nil
}


func (inst*pea8e092494_s3_Com3x) getContext(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}



// type pea8e09249.Com3z in package:github.com/starter-go/configen/src/example/parts/s3
//
// id:com-ea8e092494300fad-s3-Com3z
// class:class-e8a3937f481a2a4fcb65cb9f0011b311-IDao
// alias:
// scope:singleton
//
type pea8e092494_s3_Com3z struct {
}

func (inst* pea8e092494_s3_Com3z) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ea8e092494300fad-s3-Com3z"
	r.Classes = "class-e8a3937f481a2a4fcb65cb9f0011b311-IDao"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pea8e092494_s3_Com3z) new() any {
    return &pea8e09249.Com3z{}
}

func (inst* pea8e092494_s3_Com3z) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pea8e09249.Com3z)
	nop(ie, com)

	


    return nil
}


