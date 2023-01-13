package main

const (
	HEIGHT       = int32(768)
	WINDOW_TITLE = "3D Tic Tac Toe"
)

var (
	height = HEIGHT
	width  = (height / 9) * 16

	workingDir  string

	IsExecuting bool
)

type ()
