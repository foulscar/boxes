package main

import (
	b  "github.com/foulscar/boxes"
)


func main() {
	println("he1")
	engine := b.LoadEngine()
	println("he2")
	engine.LoadMaterials("./testing/materials/defaultMaterials.mtl")
	engine.LoadObjectsFromPointerFile("./testing/objects/objects.csv")

	engine.Run()
}
