package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Game contains all resources needed by the game
type Game struct {
	Width  int
	Height int
	snake  *Snake
	stage  *Stage
	food   *Food
	screen *bufio.Writer
	sig    chan os.Signal // listen ctrl+c
	input  chan byte      // listen keyboard press
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
	sig, input := keyPressEvent()

	game := Game{
		Width:  width,
		Height: height,
		screen: bufio.NewWriter(os.Stdout),
		sig:    sig,
		input:  input,
	}

	game.snake = NewSnake(2, 2, CharSnakeBody, game.input)
	game.food = NewFood(0, width-1, 0, height-1, game.snake.getCoords())
	game.stage = NewStage(width, height)

	return &game
}

func (game *Game) isFull() bool {
	return len(game.snake.body) >= (game.stage.width-2)*(game.stage.height-2)
}

// Run run the game
func (game *Game) Run() {
	ticker := time.NewTicker(200 * time.Microsecond)

	for range ticker.C {
		if game.snake.IsBiteSelf() || game.isFull() {
			fmt.Println("Game over, Your score is ", len(game.snake.body)-1)
		}
		game.snake.Move(game.Height, game.Width)
	}
}
