// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cboot

import (
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
	boot0x2acd90 "github.com/bitwormhole/wpm-wails/app/boot"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-boot0x2acd90.Loader
	cominfobuilder.Next()
	cominfobuilder.ID("com0-boot0x2acd90.Loader").Class("server-loader life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLoader{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLoader : the factory of component: com0-boot0x2acd90.Loader
type comFactory4pComLoader struct {

    mPrototype * boot0x2acd90.Loader

	
	mPortSelector config.InjectionSelector
	mWithGuiSelector config.InjectionSelector
	mWithServerSelector config.InjectionSelector
	mACSelector config.InjectionSelector

}

func (inst * comFactory4pComLoader) init() application.ComponentFactory {

	
	inst.mPortSelector = config.NewInjectionSelector("${server.port}",nil)
	inst.mWithGuiSelector = config.NewInjectionSelector("${wpm.options.run-with-gui}",nil)
	inst.mWithServerSelector = config.NewInjectionSelector("${wpm.options.run-with-server}",nil)
	inst.mACSelector = config.NewInjectionSelector("context",nil)


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLoader) newObject() * boot0x2acd90.Loader {
	return & boot0x2acd90.Loader {}
}

func (inst * comFactory4pComLoader) castObject(instance application.ComponentInstance) * boot0x2acd90.Loader {
	return instance.Get().(*boot0x2acd90.Loader)
}

func (inst * comFactory4pComLoader) GetPrototype() lang.Object {
	return inst.mPrototype
}

func (inst * comFactory4pComLoader) NewInstance() application.ComponentInstance {
	return config.SimpleInstance(inst, inst.newObject())
}

func (inst * comFactory4pComLoader) AfterService() application.ComponentAfterService {
	return inst
}

func (inst * comFactory4pComLoader) Init(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLoader) Destroy(instance application.ComponentInstance) error {
	return nil
}

func (inst * comFactory4pComLoader) Inject(instance application.ComponentInstance, context application.InstanceContext) error {
	
	obj := inst.castObject(instance)
	obj.Port = inst.getterForFieldPortSelector(context)
	obj.WithGui = inst.getterForFieldWithGuiSelector(context)
	obj.WithServer = inst.getterForFieldWithServerSelector(context)
	obj.AC = inst.getterForFieldACSelector(context)
	return context.LastError()
}

//getterForFieldPortSelector
func (inst * comFactory4pComLoader) getterForFieldPortSelector (context application.InstanceContext) int {
    return inst.mPortSelector.GetInt(context)
}

//getterForFieldWithGuiSelector
func (inst * comFactory4pComLoader) getterForFieldWithGuiSelector (context application.InstanceContext) bool {
    return inst.mWithGuiSelector.GetBool(context)
}

//getterForFieldWithServerSelector
func (inst * comFactory4pComLoader) getterForFieldWithServerSelector (context application.InstanceContext) bool {
    return inst.mWithServerSelector.GetBool(context)
}

//getterForFieldACSelector
func (inst * comFactory4pComLoader) getterForFieldACSelector (context application.InstanceContext) application.Context {
    return context.Context()
}




