// (todo:gen2.template) 
// 这个配置文件是由 starter-configen 工具自动生成的。
// 任何时候，都不要手工修改这里面的内容！！！

package cserver

import (
	application "github.com/bitwormhole/starter/application"
	config "github.com/bitwormhole/starter/application/config"
	lang "github.com/bitwormhole/starter/lang"
	util "github.com/bitwormhole/starter/util"
	server0x9665de "github.com/bitwormhole/wpm-wails/app/server"
    
)


func nop(x ... interface{}){
	util.Int64ToTime(0)
	lang.CreateReleasePool()
}


func autoGenConfig(cb application.ConfigBuilder) error {

	var err error = nil
	cominfobuilder := config.ComInfo()
	nop(err,cominfobuilder)

	// component: com0-server0x9665de.Loader
	cominfobuilder.Next()
	cominfobuilder.ID("com0-server0x9665de.Loader").Class("server-loader life").Aliases("").Scope("")
	cominfobuilder.Factory((&comFactory4pComLoader{}).init())
	err = cominfobuilder.CreateTo(cb)
	if err != nil {
		return err
	}



    return nil
}

////////////////////////////////////////////////////////////////////////////////

// comFactory4pComLoader : the factory of component: com0-server0x9665de.Loader
type comFactory4pComLoader struct {

    mPrototype * server0x9665de.Loader

	

}

func (inst * comFactory4pComLoader) init() application.ComponentFactory {

	


	inst.mPrototype = inst.newObject()
    return inst
}

func (inst * comFactory4pComLoader) newObject() * server0x9665de.Loader {
	return & server0x9665de.Loader {}
}

func (inst * comFactory4pComLoader) castObject(instance application.ComponentInstance) * server0x9665de.Loader {
	return instance.Get().(*server0x9665de.Loader)
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
	return nil
}




