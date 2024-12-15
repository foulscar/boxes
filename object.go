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
	any string `csv:",any"`
}

func (e *Engine) LoadObjectsFromPointerFile(pointerFilePath string) {
	bytes, err := os.ReadFile(pointerFilePath)	
	if err != nil {
		log.Fatal("Could not read object pointer file: '", pointerFilePath, "': ", err)
	}
	
	fileEntries := make([]objectPointerFileEntry, 0)
	if err := csv.Unmarshal(bytes, &fileEntries); err != nil {
		log.Fatal("'", pointerFilePath, "' is an invalid object pointer file: ", err)
	}
	for _, entry := range fileEntries {
		println(entry.any)
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
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal("Could not read object file: '", filepath, "': ", err)
	}

	boxDefs := make([]objectFileBoxDefinition, 0)
	if err := csv.Unmarshal(bytes, &boxDefs); err != nil {
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
		matrix = rl.MatrixMultiply(matrix, rl.MatrixRotateXYZ(rl.NewVector3(boxDef.rotX,boxDef.rotY,boxDef.rotZ)))

		obj.Boxes[i] = ObjectBox{
			Matrix:   matrix,
			Material: &mat,
		}
	}

	e.ResourceManager.Objects[objectID] = obj
}
