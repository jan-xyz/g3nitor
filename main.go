package main

import (
	"github.com/g3n/engine/util/application"
	"github.com/jan-xyz/g3nitor/scene"
)

func main() {

	app, _ := application.Create(application.Options{
		Title:      "g3nitor",
		Width:      800,
		Height:     600,
		Fullscreen: true,
	})

	scene.AddCube(app)
	scene.AddLighting(app)

	app.CameraPersp().SetPosition(0, 0, 3)
	app.Run()
}
