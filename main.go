package main

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/light"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
)

func main() {

	app, _ := application.Create(application.Options{
		Title:      "g3nitor",
		Width:      800,
		Height:     600,
		Fullscreen: true,
	})

	// Create a geometry and add it to the scene
	geom := geometry.NewCube(1)
	mat := material.NewPhong(math32.NewColor("Green"))
	torusMesh := graphic.NewMesh(geom, mat)
	app.Scene().Add(torusMesh)

	app.Subscribe(application.OnBeforeRender, func(evname string, ev interface{}) {
		delta := app.FrameDeltaSeconds() * 2 * math32.Pi / 10
		torusMesh.RotateY(delta)
	})

	// Add lights to the scene
	ambientLight := light.NewAmbient(&math32.Color{R: 1.0, G: 1.0, B: 1.0}, 0.8)
	app.Scene().Add(ambientLight)
	pointLight := light.NewPoint(&math32.Color{R: 1, G: 1, B: 1}, 5.0)
	pointLight.SetPosition(1, 0, 2)
	app.Scene().Add(pointLight)

	// Add an axis helper to the scene
	axis := graphic.NewAxisHelper(1.0)
	app.Scene().Add(axis)

	app.CameraPersp().SetPosition(0, 0, 3)
	app.Run()
}
