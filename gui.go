package main

import (
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

type myGUI struct {
	context *myContext
}

func (inst *myGUI) Run() error {

	ctx := inst.context
	icon, err := inst.loadIcon()
	if err != nil {
		return err
	}

	// Create an instance of the app structure
	app := NewApp(ctx)

	// Create application with options
	err = wails.Run(&options.App{
		Title:     "WPM",
		Width:     1024,
		Height:    768,
		MaxWidth:  4000,
		MaxHeight: 2500,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 11, G: 22, B: 33, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Linux: &linux.Options{
			Icon: icon,
		},
	})

	// if err != nil {
	// 	println("Error:", err.Error())
	// }

	return err
}

func (inst *myGUI) loadIcon() ([]byte, error) {
	return theIconData, nil
}
