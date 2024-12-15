package main

import (
	b  "github.com/foulscar/boxes"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func runtimeHandler(e *b.Engine) {
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(0, 2, 0)
	camera.Target   = rl.NewVector3(0, 0, 0)
	camera.Up       = rl.NewVector3(0, 1, 0)
	camera.Fovy     = 60
	camera.Projection = rl.CameraPerspective

	mainScene := b.NewScene(camera)
	for _, obj := range e.ResourceManager.Objects {
		e.InstiantiateObjectInScene(&mainScene, &obj)
	}

	editor := editor{
		scene: &mainScene,
		gridYLevel: 0,
	}

	for {
		rl.BeginDrawing()
		rl.ClearBackground(rl.SkyBlue)
		rl.BeginMode3D(camera)

		editor.drawGrid(7, rl.Black, rl.Red)
		e.DrawScene(&mainScene)

		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
