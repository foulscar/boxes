package boxes

import (
	"log"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
	csv "github.com/trimmer-io/go-csv"
)

type ObjectBox struct {
	Matrix   rl.Matrix
	Material *rl.Material
}

type Object struct {
	Boxes []ObjectBox
}

type InstancedObject struct {
	Object *Object
	Matrix rl.Matrix
}

type objectFileBoxDefinition struct {
	MatID  string  `csv:"matID"`
	PosX   float32 `csv:"posX"`
	PosY   float32 `csv:"posY"`
	PosZ   float32 `csv:"posZ"`
	ScaleX float32 `csv:"scaleX"`
	ScaleY float32 `csv:"scaleY"`
	ScaleZ float32 `csv:"scaleZ"`
	RotX   float32 `csv:"rotX"`
	RotY   float32 `csv:"rotY"`
	RotZ   float32 `csv:"rotZ"`
}

type objectPointerFileEntry struct {
	ObjectID string `csv:"objID"`
	FilePath string `csv:"file"`
}

func (e *Engine) LoadObjectsFromPointerFile(pointerFilePath string) {
	bytes, err := os.ReadFile(pointerFilePath)
	if err != nil {
		log.Fatal("Could not read object pointer file: '", pointerFilePath, "': ", err)
	}

	fileEntries := make([]*objectPointerFileEntry, 0)
	if err := csv.Unmarshal(bytes, &fileEntries); err != nil {
		log.Fatal("'", pointerFilePath, "' is an invalid object pointer file: ", err)
	}
	for _, entry := range fileEntries {
		if entry.ObjectID == "" || entry.FilePath == "" {
			log.Fatal("'", pointerFilePath, "' is an invalid object pointer file")
		}
		e.LoadObjectFile(path.Join(path.Dir(pointerFilePath), entry.FilePath), entry.ObjectID)
	}
}

func (e *Engine) LoadObjectFile(filepath, objectID string) {
	if e.ResourceManager.Objects == nil {
		log.Fatal("The Resource Manager has not been initialized")
	}
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Could not read object file: '", filepath, "': ", err)
	}

	boxDefs := make([]*objectFileBoxDefinition, 0)
	if err := csv.Unmarshal(bytes, &boxDefs); err != nil {
		log.Fatal("'", filepath, "' is an invalid object file: ", err)
	}

	obj := Object{
		Boxes: make([]ObjectBox, len(boxDefs)),
	}

	for i, boxDef := range boxDefs {
		mat, exists := e.ResourceManager.Materials[boxDef.MatID]
		if !exists {
			log.Fatal("box[", i, "] in '", filepath, "' contains an undefined matID: '", boxDef.MatID, "'")
		}
		matrix := e.ResourceManager.IdentityMatrix

		matrix = rl.MatrixMultiply(matrix, rl.MatrixTranslate(boxDef.PosX, boxDef.PosY, boxDef.PosZ))
		matrix = rl.MatrixMultiply(matrix, rl.MatrixScale(boxDef.ScaleX, boxDef.ScaleY, boxDef.ScaleZ))
		matrix = rl.MatrixMultiply(matrix, rl.MatrixRotate(rl.NewVector3(boxDef.RotX * rl.Deg2rad, boxDef.RotY * rl.Deg2rad, boxDef.RotZ * rl.Deg2rad), 1))

		obj.Boxes[i] = ObjectBox{
			Matrix:   matrix,
			Material: &mat,
		}
	}

	e.ResourceManager.Objects[objectID] = obj
}
