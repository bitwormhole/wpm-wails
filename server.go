package main

import (
	"fmt"
	"strconv"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	wpmmix "github.com/bitwormhole/wpm-mix"
	"github.com/bitwormhole/wpm-wails/app"
	"github.com/bitwormhole/wpm-wails/gen/cserver"
	"github.com/bitwormhole/wpm/server/service"
)

type myServer struct {
	context *myContext
}

// func (inst *myServer) Start() {
// 	go func() {
// 		inst.run1()
// 	}()
// }

func (inst *myServer) Run() error {
	return inst.run1()
}

func (inst *myServer) run1() error {
	go func() {
		x := recover()
		fmt.Printf("recover: %v", x)
	}()
	return inst.run2()
}

func (inst *myServer) run2() error {

	res := collection.LoadEmbedResources(&theModuleResFS, theModuleResPath)

	// module
	mb := application.ModuleBuilder{}
	mb.Name(theModuleName + "#server")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.Resources(res)
	mb.OnMount(cserver.ConfigForServer)

	mb.Dependency(wpmmix.ModuleServer())
	mod := mb.Create()

	// args := []string{}
	// args = append(args, "--server.port="+inst.getServerPort())

	i := starter.InitApp()
	i.UseMain(mod)
	service.SetAppModule(mod)

	app.BindCtx(inst.context, func(name string, c2 app.Ctx) {
		i.SetAttribute(name, c2)
	})

	// run
	return inst.runWithRuntime(i)
}

func (inst *myServer) getServerPort() string {
	port := inst.context.port
	return strconv.Itoa(port)
}

func (inst *myServer) runWithRuntime(i application.Initializer) error {
	rt, err := i.RunEx()
	if err != nil {
		return err
	}
	return rt.Loop()
}
