package main

import (
	"os"

	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
)

type myClient struct {
	context *myContext
}

func (inst *myClient) Run() {
	inst.context.port = 26666

	res := collection.LoadEmbedResources(&theModuleResFS, theModuleResPath)

	// module
	mb := application.ModuleBuilder{}
	mb.Name(theModuleName + "#client")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.Resources(res)

	mb.Dependency(starter.Module())
	mod := mb.Create()

	// args
	args := os.Args

	// run
	i := starter.InitApp()
	i.UseMain(mod)
	i.SetArguments(args)

	err := inst.runWithRuntime(i)
	if err != nil {
		panic(err)
	}
}

func (inst *myClient) runWithRuntime(i application.Initializer) error {
	rt, err := i.RunEx()
	if err != nil {
		return err
	}

	ctx := rt.Context()
	inst.loadServerPort(ctx)

	return rt.Loop()
}

func (inst *myClient) loadServerPort(ctx application.Context) {
	const name = "server.port"
	port := inst.context.port
	port = ctx.GetProperties().Getter().GetInt(name, port)
	inst.context.port = port
}
