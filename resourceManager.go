package boxes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ResourceManager struct {
	IdentityMatrix rl.Matrix
	CubeMesh  rl.Mesh
	Objects   map[string]Object
	Materials map[string]rl.Material
}

func (e *Engine) InitResourceManager() {
	if e.ResourceManager != nil {
		e.ResourceManager.Unload()
	}
	rm := ResourceManager{
		IdentityMatrix: rl.MatrixIdentity(),
		CubeMesh:  rl.GenMeshCube(1, 1, 1),
		Objects:   make(map[string]Object),
		Materials: make(map[string]rl.Material),
	}

	e.ResourceManager = &rm
}

func (rm ResourceManager) Unload() {
	rl.UnloadMesh(&rm.CubeMesh)
	for _, mat := range rm.Materials {
		rl.UnloadMaterial(mat)
	}
}
