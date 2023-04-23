package main

import "github.com/bitwormhole/wpm-wails/app"

type myContext struct {
	port int

	gui    func() error
	server func() error
	boot   func() error
}

func (inst *myContext) _Impl() app.Ctx {
	return inst
}

func (inst *myContext) GetPort() int {
	return inst.port
}

func (inst *myContext) SetPort(port int) {
	inst.port = port
}

func (inst *myContext) RunServer() error {
	fn := inst.server
	if fn == nil {
		return nil
	}
	return fn()
}

func (inst *myContext) RunGUI() error {
	fn := inst.gui
	if fn == nil {
		return nil
	}
	return fn()
}
