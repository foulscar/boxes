package boxes

import (
	"log"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
	csv "github.com/gocarina/gocsv"
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
	matID  string  `csv:"matID"`
	posX   float32 `csv:"posX"`
	posY   float32 `csv:"posY"`
	posZ   float32 `csv:"posZ"`
	scaleX float32 `csv:"scaleX"`
	scaleY float32 `csv:"scaleY"`
	scaleZ float32 `csv:"scaleZ"`
	rotX   float32 `csv:"rotX"`
	rotY   float32 `csv:"rotY"`
	rotZ   float32 `csv:"rotZ"`
}

type objectPointerFileEntry struct {
	objectID string `csv:"objID"`
	filePath string `csv:"file"`
}

func (e *Engine) LoadObjectsFromPointerFile(pointerFilePath string) {
	ptrFile, err := os.OpenFile(pointerFilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal("Could not open object pointer file: '", pointerFilePath, "': ", err)
	}
	defer ptrFile.Close()

	fileEntries := []*objectPointerFileEntry{}
	if err := csv.UnmarshalFile(ptrFile, &fileEntries); err != nil {
		log.Fatal("'", pointerFilePath, "' is an invalid object pointer file: ", err)
	}
	for _, entry := range fileEntries {
		if entry.objectID == "" || entry.filePath == "" {
			log.Fatal("'", pointerFilePath, "' is an invalid object pointer file")
		}
		e.LoadObjectFile(path.Join(path.Dir(pointerFilePath), entry.filePath), entry.objectID)
	}
}

func (e *Engine) LoadObjectFile(filepath, objectID string) {
	if e.ResourceManager.Objects == nil {
		log.Fatal("The Resource Manager has not been initialized")
	}
	objFile, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal("Could not open object file: '", filepath, "': ", err)
	}
	defer objFile.Close()

	boxDefs := []*objectFileBoxDefinition{}
	if err := csv.UnmarshalFile(objFile, &boxDefs); err != nil {
		log.Fatal("'", filepath, "' is an invalid object file: ", err)
	}

	obj := Object{
		Boxes: make([]ObjectBox, len(boxDefs)),
	}

	for i, boxDef := range boxDefs {
		mat, exists := e.ResourceManager.Materials[boxDef.matID]
		if !exists {
			log.Fatal("box[", i, "] in '", filepath, "' contains an undefined matID: '", boxDef.matID, "'")
		}
		matrix := e.ResourceManager.IdentityMatrix

		matrix = rl.MatrixMultiply(matrix, rl.MatrixTranslate(boxDef.posX, boxDef.posY, boxDef.posZ))
		matrix = rl.MatrixMultiply(matrix, rl.MatrixScale(boxDef.scaleX, boxDef.scaleY, boxDef.scaleZ))
		matrix = rl.MatrixMultiply(matrix, rl.MatrixRotateXYZ(rl.NewVector3(boxDef.rotX, boxDef.rotY, boxDef.rotZ)))

		obj.Boxes[i] = ObjectBox{
			Matrix:   matrix,
			Material: &mat,
		}
	}

	e.ResourceManager.Objects[objectID] = obj
}
