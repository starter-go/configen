package gen
import (
    p4b3a6218b "bytes"
    pf98ed07a4 "io"
    p0ef6f2938 "github.com/starter-go/application"
    pe8a3937f4 "github.com/starter-go/configen/src/example/parts"
    pcadc8c8db "sort"
    p8bcf66297 "strings"
     "github.com/starter-go/application/components"
)

// type pe8a3937f4.Com1 in package:github.com/starter-go/configen/src/example/parts
//
// id:com1-1
// class:com1 com-cadc8c8db42409733582cb3e2298ef87-Interface class-5c18ef72771564b7f43c497dc507aeab-Context
// alias:com-cadc8c8db42409733582cb3e2298ef87-Interface
// scope:singleton
//
type pe8a3937f48_parts_Com1 struct {
}

func (inst* pe8a3937f48_parts_Com1) register(cr components.Registry) {
	r := cr.New()
	r.ID = "com1-1"
	r.Classes = "com1 com-cadc8c8db42409733582cb3e2298ef87-Interface class-5c18ef72771564b7f43c497dc507aeab-Context"
	r.Aliases = "com-cadc8c8db42409733582cb3e2298ef87-Interface"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	cr.Register(r)
}

func (inst* pe8a3937f48_parts_Com1) new() any {
    return &pe8a3937f4.Com1{}
}

func (inst* pe8a3937f48_parts_Com1) inject(injection components.Injection, instance any) error {
	ie := injection.Ext()
	com := instance.(*pe8a3937f4.Com1)

	    com.Field1 = inst.getField1(ie)
    com.Field2 = inst.getField2(ie)
    com.Field3 = inst.getField3(ie)
    com.Field4 = inst.getField4(ie)
    com.Field5 = inst.getField5(ie)
    com.Field6 = inst.getField6(ie)
    com.Field7 = inst.getField7(ie)


    return nil
}


func (inst*pe8a3937f48_parts_Com1) getField1(ie components.InjectionExt)[]any{
    dst := make([]any, 0)
    src := ie.ListComponents(".class-native-any")
    for _, item1 := range src {
        item2 := item1.(any)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pe8a3937f48_parts_Com1) getField2(ie components.InjectionExt)[]pcadc8c8db.Interface{
    dst := make([]pcadc8c8db.Interface, 0)
    src := ie.ListComponents(".class-cadc8c8db42409733582cb3e2298ef87-Interface")
    for _, item1 := range src {
        item2 := item1.(pcadc8c8db.Interface)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pe8a3937f48_parts_Com1) getField3(ie components.InjectionExt)pcadc8c8db.Interface{
return ie.GetComponent("#id-cadc8c8db42409733582cb3e2298ef87-Interface").(pcadc8c8db.Interface)}


func (inst*pe8a3937f48_parts_Com1) getField4(ie components.InjectionExt)*p8bcf66297.Builder{
return ie.GetComponent("#id-8bcf6629759bd278a5c6266bd9c054f8-Builder").(*p8bcf66297.Builder)}


func (inst*pe8a3937f48_parts_Com1) getField5(ie components.InjectionExt)*p4b3a6218b.Buffer{
return ie.GetComponent("#id-4b3a6218bb3e3a7303e8a171a60fcf92-Buffer").(*p4b3a6218b.Buffer)}


func (inst*pe8a3937f48_parts_Com1) getField6(ie components.InjectionExt)pf98ed07a4.Reader{
return ie.GetComponent("#id-f98ed07a4d5f50f7de1410d905f1477f-Reader").(pf98ed07a4.Reader)}


func (inst*pe8a3937f48_parts_Com1) getField7(ie components.InjectionExt)p0ef6f2938.Context{
return ie.GetComponent("context").(p0ef6f2938.Context)}



// type pe8a3937f4.Com2 in package:github.com/starter-go/configen/src/example/parts
//
// id:com-e8a3937f481a2a4f-parts-Com2
// class:class-f98ed07a4d5f50f7de1410d905f1477f-Writer com-f98ed07a4d5f50f7de1410d905f1477f-Reader
// alias:com-f98ed07a4d5f50f7de1410d905f1477f-Reader
// scope:singleton
//
type pe8a3937f48_parts_Com2 struct {
}

func (inst* pe8a3937f48_parts_Com2) register(cr components.Registry) {
	r := cr.New()
	r.ID = "com-e8a3937f481a2a4f-parts-Com2"
	r.Classes = "class-f98ed07a4d5f50f7de1410d905f1477f-Writer com-f98ed07a4d5f50f7de1410d905f1477f-Reader"
	r.Aliases = "com-f98ed07a4d5f50f7de1410d905f1477f-Reader"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	cr.Register(r)
}

func (inst* pe8a3937f48_parts_Com2) new() any {
    return &pe8a3937f4.Com2{}
}

func (inst* pe8a3937f48_parts_Com2) inject(injection components.Injection, instance any) error {
	ie := injection.Ext()
	com := instance.(*pe8a3937f4.Com2)

	    com.F11 = inst.getF11(ie)
    com.F12 = inst.getF12(ie)
    com.F13 = inst.getF13(ie)
    com.F14 = inst.getF14(ie)
    com.F15 = inst.getF15(ie)
    com.F21 = inst.getF21(ie)
    com.F22 = inst.getF22(ie)
    com.F23 = inst.getF23(ie)
    com.F24 = inst.getF24(ie)
    com.F25 = inst.getF25(ie)
    com.F31 = inst.getF31(ie)
    com.F32 = inst.getF32(ie)
    com.F33 = inst.getF33(ie)
    com.F34 = inst.getF34(ie)
    com.F35 = inst.getF35(ie)
    com.F36 = inst.getF36(ie)
    com.F37 = inst.getF37(ie)


    return nil
}


func (inst*pe8a3937f48_parts_Com2) getF11(ie components.InjectionExt)int{
    return ie.GetInt("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF12(ie components.InjectionExt)int8{
    return ie.GetInt8("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF13(ie components.InjectionExt)int16{
    return ie.GetInt16("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF14(ie components.InjectionExt)int32{
    return ie.GetInt32("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF15(ie components.InjectionExt)int64{
    return ie.GetInt64("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF21(ie components.InjectionExt)uint{
    return ie.GetUint("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF22(ie components.InjectionExt)uint8{
    return ie.GetUint8("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF23(ie components.InjectionExt)uint16{
    return ie.GetUint16("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF24(ie components.InjectionExt)uint32{
    return ie.GetUint32("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF25(ie components.InjectionExt)uint64{
    return ie.GetUint64("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF31(ie components.InjectionExt)string{
    return ie.GetString("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF32(ie components.InjectionExt)byte{
    return ie.GetByte("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF33(ie components.InjectionExt)bool{
    return ie.GetBool("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF34(ie components.InjectionExt)rune{
    return ie.GetRune("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF35(ie components.InjectionExt)float32{
    return ie.GetFloat32("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF36(ie components.InjectionExt)float64{
    return ie.GetFloat64("${a.b.c.d}")
}


func (inst*pe8a3937f48_parts_Com2) getF37(ie components.InjectionExt)any{
    return ie.GetAny("${a.b.c.d}")
}



// type pe8a3937f4.Com3 in package:github.com/starter-go/configen/src/example/parts
//
// id:com-3-abc
// class:
// alias:c3-xyz c3-ijk
// scope:singleton
//
type pe8a3937f48_parts_Com3 struct {
}

func (inst* pe8a3937f48_parts_Com3) register(cr components.Registry) {
	r := cr.New()
	r.ID = "com-3-abc"
	r.Classes = ""
	r.Aliases = "c3-xyz c3-ijk"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	cr.Register(r)
}

func (inst* pe8a3937f48_parts_Com3) new() any {
    return &pe8a3937f4.Com3{}
}

func (inst* pe8a3937f48_parts_Com3) inject(injection components.Injection, instance any) error {
	ie := injection.Ext()
	com := instance.(*pe8a3937f4.Com3)

	    com.A = inst.getA(ie)


    return nil
}


func (inst*pe8a3937f48_parts_Com3) getA(ie components.InjectionExt)int{
    return ie.GetInt("#id-native-int")
}


