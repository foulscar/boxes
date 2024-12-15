package main

import (
        "math"
        rl "github.com/gen2brain/raylib-go/raylib"
)

func (scn *scene) drawGrid(numLines float32, color rl.Color, color2 rl.Color) {
	pos := rl.NewVector3(
		float32(math.Round(float64(scn.camera.Position.X))),
		float32(scn.gridYLevel),
		float32(math.Round(float64(scn.camera.Position.Z))),
	)
	for i := -numLines - 2.0; i <= numLines+2.0; i++ {
		rl.DrawLine3D(rl.NewVector3(pos.X-numLines, pos.Y, pos.Z+i), rl.NewVector3(pos.X+numLines, pos.Y, pos.Z+i), color)
		rl.DrawLine3D(rl.NewVector3(pos.X+i, pos.Y, pos.Z-numLines), rl.NewVector3(pos.X+i, pos.Y, pos.Z+numLines), color)
	}
	factors := []float32{-1.0, 1.0}
	for _, xFactor := range factors {
		for _, zFactor := range factors {
			rl.DrawLine3D(
				rl.NewVector3(pos.X+((numLines+1.0)*xFactor), pos.Y, pos.Z+((numLines+1.0)*zFactor)),
				rl.NewVector3(pos.X+((numLines+2.0)*xFactor), pos.Y, pos.Z+((numLines+2.0)*zFactor)),
				color2,
			)
			rl.DrawLine3D(
				rl.NewVector3(pos.X+((numLines+2.0)*xFactor), pos.Y, pos.Z+((numLines+1.0)*zFactor)),
				rl.NewVector3(pos.X+((numLines+1.0)*xFactor), pos.Y, pos.Z+((numLines+2.0)*zFactor)),
				color2,
			)
		}
	}
}
