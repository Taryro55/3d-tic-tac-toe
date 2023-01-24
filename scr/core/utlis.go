// Variated utility functions, like conversions, contains, isX, ect

package core

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func InvertNumb(numb float32) float32 {
	return numb * -1
}

func IsBetween(x int, a, b int) bool {
	if a <= x && x <= b {
		return true
	} else {
		return false
	}
}

func (block *Blocks) SetChunksPosition() {
	var p int
	for i := -1; i < 2; i++ {
		var height = i
		for x := 0; x < (ChunkAmmount)/3; x++ {
			var pos rl.Vector3

			disperser := rl.NewVector3(
				Board.Spacing.X*(1-float32(2*x)),
				Board.Spacing.Y*float32(height),
				Board.Spacing.Z*(-2*float32(x)+5),
			)

			val := Board.Spacing.X * (float32(math.Sin(rl.Pi*float64(x) + rl.Pi/2)))
			disperserDiag := rl.NewVector3(
				val,
				Board.Spacing.Y*float32(height),
				InvertNumb(val),
			)

			switch {
			case IsBetween(x, 0, 1):
				pos = rl.NewVector3(disperser.X, disperser.Y, 0.0)
			case IsBetween(x, 2, 3):
				pos = rl.NewVector3(0.0, disperser.Y, disperser.Z)
			case IsBetween(x, 4, 5):
				pos = rl.NewVector3(disperserDiag.X, disperserDiag.Y, disperserDiag.Z)
			case IsBetween(x, 6, 7):
				pos = rl.NewVector3(disperserDiag.X, disperserDiag.Y, disperserDiag.Z*-1)
			default:
				pos = rl.NewVector3(0.0, disperserDiag.Y, 0.0)
			}

			block.Chunks[p].Position = pos
			p++
		}
	}
}

func (block *Blocks) SetChunksColors() {
	for i := range block.Chunks {
		block.Chunks[i].ColorMain = rl.Red
		block.Chunks[i].ColorSecond = rl.Maroon
	}
}

func (block *Blocks) SetChunksEdge() {
	for i := range block.Chunks {
		block.Chunks[i].Edge = EDGE
	}
}
