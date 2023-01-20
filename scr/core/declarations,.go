// General Variables

package core

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	FPS          = 60
	HEIGHT       = int32(768)
	WINDOW_TITLE = "3D Tic Tac Toe!"
)

var (
	Height = HEIGHT
	Width  = (Height / 9) * 16

	workingDir string

	IsExecuting bool
)

var (
	CenterPos = rl.NewVector3(0.0, 0.0, 0.0)
	Camera = rl.Camera3D{
		Position:   rl.NewVector3(10.0, 10.0, 10.0),
		Target:     rl.NewVector3(0.0, 0.0, 0.0),
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:       25.0,
		Projection: rl.CameraOrthographic,
	}

	Board = Blocks{
		rl.Red,
		27,
		make([][][]int, 0),
		false,
		rl.NewVector3(4.0, 4.0, 4.0),
		CenterPos,
		rl.Red,
		rl.Maroon,
	}

	//
)

type (
	Blocks struct {
		Color           rl.Color
		ChunkCount      int
		Chunks          [][][]int
		IsDeconstructed bool
		Spacing			rl.Vector3
		Position        rl.Vector3
		MainColor   rl.Color
		SecondColor rl.Color
	}
	Game struct {
		IsOver bool
		Length float64
		Winner bool
	}

	Menu struct {
		IsActive    bool
		Background  rl.Color
		MainColor   rl.Color
		SecondColor rl.Color
	}
)
