package main

import (
	// "os"

	"tictactoe-3d/scr/gui"
	"tictactoe-3d/scr/logic"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func init() {
	// workingDir, _ := os.Getwd()

	rl.InitWindow(width, height, WINDOW_TITLE)
	rl.SetTargetFPS(60)
	rl.SetMouseScale(1.0, 1.0)
}

func quit() {
	rl.CloseWindow()
}

func main() {
	for IsExecuting {
		logic.Update()
		gui.Render()
	}
	quit()
}
