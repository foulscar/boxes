package main

import (
	b  "github.com/foulscar/boxes"
)


func main() {
	engine := b.LoadEngine()
	engine.LoadMaterials("./testing/materials/defaultMaterials.mtl")
	engine.LoadObjectsFromPointerFile("./testing/objects/objects.csv")
}
