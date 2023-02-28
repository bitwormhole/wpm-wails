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
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.Resources(res)

	mb.Dependency(starter.Module())
	mod := mb.Create()

	// args
	args := []string{}
	args = os.Args

	// run
	i := starter.InitApp()
	i.UseMain(mod)
	i.SetArguments(args)
	i.Run()
}
