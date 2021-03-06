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
	debug       bool
	screen      *bufio.Writer
	sig         chan os.Signal // listen ctrl+c
	directionCh chan Direction // listen keyboard press
}

// GameOpts is used to init the game
type GameOpts struct {
	Width  int
	Height int
	Debug  bool
}

// NewGame returns initialized game
func NewGame(opts GameOpts) *Game {
	width := defaultGameWidth
	if opts.Width > 2 {
		width = opts.Width
	}
	height := defaultGameHeight
	if opts.Height > 2 {
		height = opts.Height
	}
	sig, directionCh := keyPressEvent()

	game := Game{
		Width:       width,
		Height:      height,
		debug:       opts.Debug,
		screen:      bufio.NewWriter(os.Stdout),
		sig:         sig,
		directionCh: directionCh,
	}

	game.snake = NewSnake(2, 2, CharSnakeBody)

	trimWidth := width - 2
	trimHeight := height - 2
	game.food = NewFood(1, trimWidth, 1, trimHeight, game.snake.getCoords())
	game.stage = NewStage(width, height)

	return &game
}

func (game *Game) clear() {
	game.screen.Write(CharClear)
	game.screen.Flush()
}

func (game *Game) frame() []byte {
	game.snake.Move(game.Height-1, game.Width-1)
	if game.food.coord.x == game.snake.head.x && game.food.coord.y == game.snake.head.y {
		game.food.newLocate(
			1, game.stage.width-1,
			1, game.stage.height-1,
			game.snake.getCoords().concat(game.food.getCoordList()),
		)
		game.snake.body.push(game.snake.head)
	}

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
	frame := game.frame()
	if game.debug {
		game.snake.getCoords().print("snake")
		game.food.getCoordList().print("food")
	} else {
		game.clear()
		game.screen.Write(frame)
		game.screen.Flush()
	}
}

func (game *Game) isFull() bool {
	return len(*game.snake.body) >= (game.stage.width-2)*(game.stage.height-2)
}

func (game *Game) score() int {
	return len(*game.snake.body) - 1
}

// Run run the game
func (game *Game) Run() {
	nonOutputNobuffer()

	ticker := time.NewTicker(defaultSpeed)

	postCondition := func() {
		ticker.Stop()
		recoverNonOutputNobuffer()
		os.Exit(0)
	}

	for {
		select {
		case <-ticker.C:
			game.snake.apply()
			if game.snake.IsBiteSelf() || game.isFull() {
				fmt.Println("Game over, Your score is ", game.score())
				postCondition()
			}
			game.draw()
		case direction := <-game.directionCh:
			game.snake.changeDirection(direction)
		case <-game.sig:
			postCondition()
		}
	}
}
