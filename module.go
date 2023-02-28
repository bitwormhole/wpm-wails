package main

import "embed"

const (
	theModuleName     = "github.com/bitwormhole/wpm-wails"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS
