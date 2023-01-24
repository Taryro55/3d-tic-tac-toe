// Handles inputs

package logic

import (
	"tictactoe-3d/scr/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Input() {
	cameraInput(&core.Camera)
	hitbox()
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

func hitbox() {
	for i := range core.Board.Chunks {
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			chunk := &core.Board.Chunks[i]
			if !chunk.Collision.Hit {
				chunk.Ray = rl.GetMouseRay(rl.GetMousePosition(), core.Camera)
				chunk.Collision = rl.GetRayCollisionBox(
					chunk.Ray,
					rl.NewBoundingBox(
						rl.NewVector3(
							chunk.Position.X-chunk.Edge/2, chunk.Position.Y-chunk.Edge/2, chunk.Position.Z-chunk.Edge/2,
						), rl.NewVector3(
							chunk.Position.X+chunk.Edge/2, chunk.Position.Y+chunk.Edge/2, chunk.Position.Z+chunk.Edge/2,
						),
					),
				)
			} else {
				chunk.Collision.Hit = false
			}
		}
	}
}
