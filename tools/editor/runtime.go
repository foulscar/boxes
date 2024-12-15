package main

import (
	b  "github.com/foulscar/boxes"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowWidth = 1240
	windowHeight = 720
)

func runtimeHandler(e *b.Engine) {
	rl.InitWindow(windowWidth, windowHeight)
	defer rl.CloseWindow()

	e.

	for {
		rl.BeginDrawing()
		rl.BeginMode3D()

		e.Draw

		rl.EndMode3D()
		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
}
