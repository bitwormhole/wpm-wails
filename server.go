package main

import (
	"fmt"
	"strconv"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	wpmmix "github.com/bitwormhole/wpm-mix"
)

type myServer struct {
	context *myContext
}

func (inst *myServer) Start() {
	go func() {
		inst.run1()
	}()
}

func (inst *myServer) run1() {
	go func() {
		x := recover()
		fmt.Printf("recover: %v", x)
	}()
	inst.run2()
}

func (inst *myServer) run2() {

	res := collection.LoadEmbedResources(&theModuleResFS, theModuleResPath)

	// module
	mb := application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.Resources(res)

	mb.Dependency(wpmmix.Module())
	mod := mb.Create()

	args := []string{}
	args = append(args, "--server.port="+inst.getServerPort())

	// run
	i := starter.InitApp()
	i.UseMain(mod)
	i.SetArguments(args)
	i.Run()
}

func (inst *myServer) getServerPort() string {
	port := inst.context.port
	return strconv.Itoa(port)
}
