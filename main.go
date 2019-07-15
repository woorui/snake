package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	watchInterrupt(exit)

	screen := bufio.NewWriter(os.Stdout)
	stage := newStage(50, 25)
	snake := newSnake()
	// Can not locate in stage's border
	food := newFood(1, stage.width-1, 1, stage.height-1, snake.getCoords())

	input := make(chan byte)
	go keyPress(input, snake.speed)
	go snake.adapt(input)

	render(screen, stage, snake, food)
}

func render(screen *bufio.Writer, stage *stage, snake *snake, food *food) {
	ticker := time.NewTicker(snake.speed)
	for range ticker.C {
		screenClear(screen)
		if snake.checkCollidingSelf() || len(snake.body) >= (stage.width-2)*(stage.height-2) {
			fmt.Println("Game over, Your score is ", len(snake.body)-1)
			exit()
		}
		snake.move(stage, food)
		screenWrite(screen, stage.draw(append(snake.getCoords(), food.getCoords()...)))
		screenFlush(screen)
	}
}

func exit() {
	deCleanScreen()
	os.Exit(0)
}
