package logic

import (
	"tictactoe-3d/scr"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Update() {
	main.IsExecuting = !rl.WindowShouldClose()

}
