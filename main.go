package main

import (
	"embed"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	ctx := &myContext{}
	client := &myClient{context: ctx}
	server := &myServer{context: ctx}
	w := &myWails{context: ctx}

	client.Run()
	server.Start()
	w.Run()
}
