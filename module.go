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

//go:embed "src/main/resources/icon_wpm.png"
var theIconData []byte
