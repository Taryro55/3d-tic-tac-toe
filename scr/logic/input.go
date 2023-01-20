// Handles inputs

package logic

import (
	"tictactoe-3d/scr/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Input() {
	cameraInput(&core.Camera)
}

func cameraInput(camera *rl.Camera3D) {
	if rl.IsKeyDown(rl.KeyW) && camera.Position.Y < 45 {
		camera.Position.Y += 2.5
	} else if rl.IsKeyDown(rl.KeyS) && camera.Position.Y > -45 {
		camera.Position.Y -= 2.5
	} else if rl.IsKeyDown(rl.KeyA) && camera.Position.Z > -80 {
		camera.Position.Z -= 2.5
	} else if rl.IsKeyDown(rl.KeyD) && camera.Position.Z < 80 {
		camera.Position.Z += 2.5
	} else if rl.IsKeyDown(rl.KeyQ) && camera.Position.X > -80 {
		camera.Position.X -= 2.5
	} else if rl.IsKeyDown(rl.KeyE) && camera.Position.X < 80 {
		camera.Position.X += 2.5
	}

}
