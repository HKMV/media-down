package main

import (
	"embed"
	"media-down/backend"
	"media-down/backend/pkg/logs"
	wails2 "media-down/backend/pkg/wails"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := backend.NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "媒体下载",
		Width:  500,
		Height: 110,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		OnStartup:        app.Startup,
		Bind:             app.GetBinds(),
		Logger:           wails2.NewMultiLogger(logs.LogName()),
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
