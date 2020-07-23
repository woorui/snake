package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Game contains all resources needed by the game
type Game struct {
	Width       int
	Height      int
	snake       *Snake
	stage       *Stage
	food        *Food
	screen      *bufio.Writer
	sig         chan os.Signal // listen ctrl+c
	directionCh chan Direction // listen keyboard press
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
	sig, directionCh := keyPressEvent()

	game := Game{
		Width:       width,
		Height:      height,
		screen:      bufio.NewWriter(os.Stdout),
		sig:         sig,
		directionCh: directionCh,
	}

	game.snake = NewSnake(2, 2, CharSnakeBody)
	game.food = NewFood(0, width-2, 0, height-2, game.snake.getCoords())
	game.stage = NewStage(width, height)

	return &game
}

func (game *Game) clear() {
	game.screen.Write(CharClear)
	game.screen.Flush()
}

func (game *Game) stuff() []byte {
	b := make([]byte, len(game.stage.matrix))
	copy(b, game.stage.matrix)
	coords := game.snake.getCoords().concat(game.food.getCoordList())
	for _, c := range coords {
		index := game.stage.mapping[cantorPairingFn(c.x, c.y)]
		b[index] = c.ink
	}
	return append(b, CharBreaker)
}

func (game *Game) draw() {
	game.screen.Write(game.stuff())
	game.screen.Flush()
}

func (game *Game) isFull() bool {
	return len(game.snake.body) >= (game.stage.width-2)*(game.stage.height-2)
}

func (game *Game) score() int {
	return len(game.snake.body) - 1
}

// Run run the game
func (game *Game) Run() {
	nonOutputAndNobuffer()

	ticker := time.NewTicker(2000 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			game.snake.unLockDirection()
			if game.snake.IsBiteSelf() || game.isFull() {
				fmt.Println("Game over, Your score is ", game.score())
			}
			game.snake.Move(game.Height, game.Width)
			// game.draw()
			continue
		case direction := <-game.directionCh:
			game.snake.changeDirection(direction)
			continue
		case <-game.sig:
			ticker.Stop()
			recoverNonOutputAndNobuffer()
			os.Exit(0)
		}
	}
}
