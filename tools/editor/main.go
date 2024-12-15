package main

import (
	"github.com/foulscar/boxes"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth = 1240
	windowHeight = 720
)

func main() {
	rl.InitWindow(windowWidth, windowHeight, "Boxes: Testing")
	rl.SetTargetFPS(60)
	rl.SetBlendMode(rl.BlendAlpha)

	rl.DisableCursor()

	scn := newScene()

	defer rl.CloseWindow()
	for {
		if rl.IsKeyPressed(rl.KeyEscape) {
			break
		}
		if rl.IsKeyPressed(rl.KeyMinus) {
			scn.gridYLevel--
		}
		if rl.IsKeyPressed(rl.KeyEqual) {
			scn.gridYLevel++
		}
		rl.UpdateCamera(&scn.camera, rl.CameraFree)

		rl.BeginDrawing()
			rl.ClearBackground(rl.SkyBlue)
			rl.BeginMode3D(scn.camera)
				rl.BeginBlendMode(rl.BlendAlpha)
				scn.draw()
				rl.EndBlendMode()
			rl.EndMode3D()
				rl.DrawFPS(10,10)
		rl.EndDrawing()
	}
}
