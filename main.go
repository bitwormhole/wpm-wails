package main

import (
	"embed"
	"os"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	ctx := &myContext{}
	client := &myClient{context: ctx}
	// server := &myServer{context: ctx}
	w := &myWails{context: ctx}

	// vlog.Trace(client, server, w)

	x, ok := os.LookupEnv("WPM_WAILS_SKIP_RUN")
	if x == "1" && ok {
		return
	}

	client.Run()
	// server.Start()
	w.Run()
}
