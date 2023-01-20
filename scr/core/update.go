// Everything that will be executed at the start of each frame

package core

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Update() {
	IsExecuting = !rl.WindowShouldClose()
	rl.UpdateCamera(&Camera)
}

