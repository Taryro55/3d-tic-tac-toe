package main

import (
	// "os"

	"tictactoe-3d/scr/core"
	"tictactoe-3d/scr/gui"

	"tictactoe-3d/scr/logic"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func init() {
	// workingDir, _ := os.Getwd()
	core.IsExecuting = true

	rl.InitWindow(core.Width, core.Height, core.WINDOW_TITLE)
	rl.SetTargetFPS(core.FPS)
	rl.SetMouseScale(1.0, 1.0)
	rl.SetCameraMode(core.Camera, rl.CameraCustom)
}

func quit() {
	rl.CloseWindow()
}

func main() {
	core.Board.SetChunksEdge()
	core.Board.SetChunksColors()
	core.Board.SetChunksPosition()

	for core.IsExecuting {
		core.Update()
		logic.Input()
		gui.Render()

		// if rl.IsMouseButtonDown(0) {
		// 	for i := range core.Board.Chunks {
		// 		fmt.Println(core.Board.Chunks[i].Ray)
		// 	}
		// }
	}
	quit()
}
