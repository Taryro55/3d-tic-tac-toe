// Misc rendering

package gui

import (
	"math"
	"tictactoe-3d/scr/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()
	rl.ClearBackground(rl.Black)

	//
	rl.BeginMode3D(core.Camera)

	blockRendering()

	rl.DrawGrid(10, 1.0)
	rl.EndMode3D()
}

func blockRendering() {
	for i := -1; i < 2; i++ {
		var height = i
		for x := 0; x < core.Board.ChunkCount/3; x++ {
			var pos rl.Vector3
			disperser := rl.NewVector3(
				core.Board.Spacing.X*(1-float32(2*x)),
				core.Board.Spacing.Y*float32(height),
				core.Board.Spacing.Z*(-2*float32(x)+5),
			)

			var val = core.Board.Spacing.X * (float32(math.Sin(rl.Pi*float64(x) + rl.Pi/2)))
			disperserDiag := rl.NewVector3(
				val,
				core.Board.Spacing.Y*float32(height),
				core.InvertNumb(val),
			)

			if x < 2 {
				pos = rl.NewVector3(disperser.X, disperser.Y, 0.0)
			} else if 2 <= x && x <= 3 {
				pos = rl.NewVector3(0.0, disperser.Y, disperser.Z)
			} else if 4 <= x && x <= 5 {
				pos = rl.NewVector3(disperserDiag.X, disperserDiag.Y, disperserDiag.Z)
			} else if 6 <= x && x <= 7 {
				pos = rl.NewVector3(disperserDiag.X, disperserDiag.Y, disperserDiag.Z*-1)
			} else {
				pos = rl.NewVector3(0.0, disperserDiag.Y, 0.0)
			}

			rl.DrawCube(pos, 2.0, 2.0, 2.0, core.Board.MainColor)
			rl.DrawCubeWires(pos, 2.0, 2.0, 2.0, core.Board.SecondColor)
		}
	}
}
