// Misc rendering

package gui

import (
	// "math"

	"tictactoe-3d/scr/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Render() {
	rl.BeginDrawing()
	defer rl.EndDrawing()
	rl.ClearBackground(rl.Black)

	rl.BeginMode3D(core.Camera)

	blockRendering()
	rl.DrawGrid(10, 1.0)

	rl.EndMode3D()
}

func blockRendering() {
	for i, chunk := range core.Board.Chunks {
		if core.Board.Chunks[i].Collision.Hit {
			chunk.ColorMain, chunk.ColorSecond = rl.Yellow, rl.Orange
		}
		rl.DrawCube(chunk.Position, chunk.Edge, chunk.Edge, chunk.Edge, chunk.ColorMain)
		rl.DrawCubeWires(chunk.Position, chunk.Edge, chunk.Edge, chunk.Edge, chunk.ColorSecond)
	}
}
