package main

import (
	"github.com/bitwormhole/starter"
	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	wpmmix "github.com/bitwormhole/wpm-mix"
	"github.com/bitwormhole/wpm-wails/app"
	"github.com/bitwormhole/wpm-wails/gen/cboot"
)

type myBoot struct {
	context *myContext
}

func (inst *myBoot) Run() error {
	inst.context.port = 26666

	res := collection.LoadEmbedResources(&theModuleResFS, theModuleResPath)

	// module
	mb := application.ModuleBuilder{}
	mb.Name(theModuleName + "#boot")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.Resources(res)
	mb.OnMount(cboot.ConfigForBoot)

	mb.Dependency(wpmmix.ModuleBoot())
	mod := mb.Create()

	// args
	// args := os.Args

	// run
	i := starter.InitApp()
	i.UseMain(mod)
	// i.SetArguments(args)

	app.BindCtx(inst.context, func(name string, c2 app.Ctx) {
		i.SetAttribute(name, c2)
	})

	err := inst.runWithRuntime(i)
	if err != nil {
		return (err)
	}

	return nil
}

func (inst *myBoot) runWithRuntime(i application.Initializer) error {
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
