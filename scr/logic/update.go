package logic

import (
	"tictactoe-3d/scr/core"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Update() {
	core.IsExecuting = !rl.WindowShouldClose()

}
