// General Variables

package core

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	FPS          = 60
	HEIGHT       = int32(768)
	WINDOW_TITLE = "3D Tic Tac Toe!"
)

var (
	Height = HEIGHT
	Width  = (Height / 9) * 16

	workingDir string

	IsExecuting  bool
	ChunkAmmount = int(math.Pow(3, 3))

	MainColor = rl.Red
	SecondColor = rl.Maroon
)

const (
	EDGE  = 3.0
	SPACE = 4.0
	
)

var (
	CenterPos = rl.NewVector3(0.0, 0.0, 0.0)
	Camera    = rl.Camera3D{
		Position:   rl.NewVector3(10.0, 10.0, 10.0),
		Target:     rl.NewVector3(0.0, 0.0, 0.0),
		Up:         rl.NewVector3(0.0, 1.0, 0.0),
		Fovy:       25.0,
		Projection: rl.CameraOrthographic,
	}

	// Cube = Chunk{
	// 	IsSelected: false,
	// 	Owner:      0,
	// 	Position:   rl.NewVector3(0.0, 0.0, 0.0),
	// 	Edge:       3.0,
	// 	Ray:        rl.NewRay(rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector3(0.0, 0.0, 0.0)),
	// 	Collision:  rl.NewRayCollision(false, 0.0, rl.NewVector3(0.0, 0.0, 0.0), rl.NewVector3(0.0, 0.0, 0.0)),
	// }

	Board = Blocks{
		Chunks:      make([]Chunk, ChunkAmmount),
		ChunkStates: make([][][]int, 0),
		Spacing:     rl.NewVector3(SPACE, SPACE, SPACE),
		Position:    CenterPos,
	}
)

type (
	Blocks struct {
		Chunks      []Chunk
		ChunkStates [][][]int
		Spacing     rl.Vector3
		Position    rl.Vector3
	}

	Chunk struct {
		IsSelected bool
		Owner      int
		Position   rl.Vector3
		Edge       float32

		Ray       rl.Ray
		Collision rl.RayCollision

		ColorMain   rl.Color
		ColorSecond rl.Color
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
