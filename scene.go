package boxes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	Camera           rl.Camera3D
	InstancedObjects map[int]InstancedObject
}

func NewScene(camera rl.Camera3D) Scene {
	return Scene{
		Camera:           camera,
		InstancedObjects: make(map[int]InstancedObject),
	}
}

func (e *Engine) DrawScene(scn *Scene) {
	for _, instObj := range scn.InstancedObjects {
		for _, box := range instObj.Object.Boxes {
			rl.DrawMesh(
				e.ResourceManager.CubeMesh,
				*box.Material,
				rl.MatrixMultiply(instObj.Matrix, box.Matrix),
			)
		}
	}
}
