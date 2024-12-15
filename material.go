package boxes

import (
	"os"
	"log"

	rl  "github.com/gen2brain/raylib-go/raylib"
	common "github.com/mokiat/go-data-front/common"
	mtl "github.com/mokiat/go-data-front/scanner/mtl"
)

func (e *Engine) LoadMaterials(filepath string) {
	println(len(e.ResourceManager.Materials))
	if e.ResourceManager.Materials == nil {
		log.Fatal("The Resource Manager has not been initialized")
	}

	materials := rl.LoadMaterials(filepath)
	matIDs    := make([]string, 0)

	unloadOnError := func() {
		for _, mat := range materials {
			rl.UnloadMaterial(mat)
		}
	}

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		unloadOnError()
		log.Fatal("Could not open '", filepath, "' to load materials")
	}
	scanner := mtl.NewScanner()
	scanner.Scan(file, func(event common.Event) error {
		switch actual := event.(type) {
		case mtl.MaterialEvent:
			matIDs = append(matIDs, actual.MaterialName)
		}
		return nil
	})

	if len(materials) != len(matIDs) {
		unloadOnError()
		log.Fatal("Could not parse material names in '", filepath, "'")
	}

	for _ = range matIDs {
		e.ResourceManager.Materials["test"] = rl.LoadMaterialDefault()
	}
}
