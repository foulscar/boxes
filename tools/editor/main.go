package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	b  "github.com/foulscar/boxes"
)

const (
	windowWidth = 1240
	windowHeight = 720
)

func main() {
	rl.InitWindow(windowWidth, windowHeight, "Boxes Editor")
	defer rl.CloseWindow()

	engine := b.LoadEngine()

	engine.SetRuntimeHandler(runtimeHandler)
	engine.LoadMaterials("./testing/materials/defaultMaterials.mtl")
	engine.LoadObjectsFromPointerFile("./testing/objects/objects.csv")

	engine.Run()
}
