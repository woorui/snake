package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Game contains all resources needed by the game
type Game struct {
	option      *GameOption
	snake       *Snake
	food        Coordinate
	framer      Framer
	directionCh chan Direction // listen keyboard press
}

// GameOptions is used to init the game.
//
// Speed defines duration between two frames,
// The duration is time.Millisecond * GameOption.Speed.
type GameOption struct {
	Width  int
	Height int
	Speed  int
}

// NewGame returns initialized game.
func NewGame(readWriter io.ReadWriter, option *GameOption) *Game {
	var (
		height = option.Height
		width  = option.Width
	)

	var (
		minX = 1
		maxX = width - 2
		minY = 1
		maxY = height - 2
	)

	var (
		snake   = NewSnake(2, 2, CharSnakeHead, &defaultController{})
		food, _ = NewCoordinate(minX, maxX, minY, maxY, CharFood, snake.CoordinateList())
		framer  = NewFramer(readWriter, width, height)
	)

	return &Game{
		option:      option,
		snake:       snake,
		framer:      framer,
		food:        food,
		directionCh: ReadToDirection(readWriter),
	}
}

func (game *Game) PlayFrame() error {
	game.framer.Clear()
	// snake ate a food.
	if game.food.x == game.snake.head.x && game.food.y == game.snake.head.y {
		food, err := NewCoordinate(
			1, game.option.Width-1,
			1, game.option.Height-1,
			CharFood,
			Append(game.snake.CoordinateList(), game.food),
		)
		if err != nil {
			return err
		}

		game.food = food
		game.snake.body = Append(game.snake.body, game.snake.head)
	}

	return game.framer.RenderWithCoordinate(game.snake, game.food)
}

func (game *Game) score() int { return len(game.snake.body) + 1 }

// Run run the game
func (game *Game) Run() {
	close, err := OpenGameMode()
	if err != nil {
		log.Fatalln(err)
	}
	defer close()

	ticker := time.NewTicker(time.Millisecond * time.Duration(game.option.Speed))
	defer ticker.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	var (
		width  = game.option.Width
		height = game.option.Height
	)

	topScore := (width - 2) * (height - 2)

	for {
		select {
		case <-ticker.C:
			if game.score() >= topScore {
				fmt.Println("You are the King of Snake !", game.score())
				return
			}
			if game.snake.IsBumpSelf() {
				goto GAMEOVER
			}

			if err := game.PlayFrame(); err != nil {
				fmt.Println("Your snake is too fat.", game.score())
				return
			}

			// snake's scope.
			var (
				minX = 0
				maxX = width - 1
				minY = 0
				maxY = height - 1
			)

			game.snake.Move(minX, maxX, minY, maxY)

		case direction := <-game.directionCh:
			game.snake.ChangeDirection(direction)
		case <-c:
			goto GAMEOVER
		}
	}

GAMEOVER:
	fmt.Println("Game over, Your score is ", game.score())
}
