package scene

import (
	"github.com/g3n/engine/geometry"
	"github.com/g3n/engine/graphic"
	"github.com/g3n/engine/material"
	"github.com/g3n/engine/math32"
	"github.com/g3n/engine/util/application"
)

// AddCube adds a cube to an scene
func AddCube(app *application.Application) {

	// Create a geometry and add it to the scene
	geom := geometry.NewCube(1)
	mat := material.NewPhong(math32.NewColor("Green"))
	torusMesh := graphic.NewMesh(geom, mat)
	app.Scene().Add(torusMesh)

	app.Subscribe(application.OnBeforeRender, func(evname string, ev interface{}) {
		delta := app.FrameDeltaSeconds() * 2 * math32.Pi / 10
		torusMesh.RotateY(delta)
	})
}
