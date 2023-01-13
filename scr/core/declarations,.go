package core

const (
	HEIGHT       = int32(768)
	WINDOW_TITLE = "3D Tic Tac Toe"
)

var (
	Height = HEIGHT
	Width  = (Height / 9) * 16

	workingDir  string

	IsExecuting bool
)

type ()
