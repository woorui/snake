package main

import (
	"bufio"
	"os"
)

// Game contains all resources needed by the game
type Game struct {
	Width  int8
	Height int8
	snake  snake
	stage  stage
	food   food
	screen *bufio.Writer
}

// GameOpts is used to init the game
type GameOpts struct {
	Width  int
	Height int
}

// NewGame returns initialized game
func NewGame(opts GameOpts) *Game {
	width := defaultGameWidth
	if opts.Width != 0 {
		width = opts.Width
	}
	height := defaultGameHeight
	if opts.Height != 0 {
		height = opts.Height
	}

	return &Game{
		Width:  int8(width),
		Height: int8(height),
		screen: bufio.NewWriter(os.Stdout),
	}
}
