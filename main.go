package main

import (
	"embed"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	ctx := &myContext{}
	boot := &myBoot{context: ctx}
	gui := &myGUI{context: ctx}
	server := &myServer{context: ctx}

	ctx.boot = boot.Run
	ctx.gui = gui.Run
	ctx.server = server.Run

	err := ctx.boot()
	if err != nil {
		panic(err)
	}
}
