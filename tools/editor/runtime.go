package main

import (
	b  "github.com/foulscar/boxes"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func runtimeHandler(e *b.Engine) {
	camera := rl.NewCamera3D(
		rl.NewVector3(1, 10, 1),
		rl.NewVector3(0, 0, 0),
		rl.NewVector3(0, 1, 0),
		60,
		rl.CameraPerspective,
	)

	mainScene := b.NewScene(&camera)
	for _, obj := range e.ResourceManager.Objects {
		e.InstiantiateObjectInScene(&mainScene, &obj)
	}

	editor := editor{
		gridYLevel: 0,
		scene: &mainScene,
	}

	rl.DisableCursor()
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree)

		if rl.IsKeyPressed(rl.KeyMinus) {
			editor.gridYLevel--
		}

		if rl.IsKeyPressed(rl.KeyEqual) {
			editor.gridYLevel++
		}

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
