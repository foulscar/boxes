package boxes

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	Camera           *rl.Camera3D
	InstancedObjects map[int]InstancedObject
}

func NewScene(camera *rl.Camera3D) Scene {
	return Scene{
		Camera:           camera,
		InstancedObjects: make(map[int]InstancedObject),
	}
}

func (e *Engine) InstiantiateObjectInScene(scn *Scene, obj *Object) (instanceID int) {
	randID := 0
	for {
		randID = rand.Int()
		if _, exists := scn.InstancedObjects[randID]; !exists {
			break
		}
	}

	scn.InstancedObjects[randID] = InstancedObject{
		Object: obj,
		Matrix: e.ResourceManager.IdentityMatrix,
	}

	return randID
}

func (e *Engine) RemoveInstanceFromScene(scn *Scene, id int) {
	delete(scn.InstancedObjects, id)
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
