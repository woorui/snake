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
	go keyPress(input)
	go snake.adapt(input)

	render(screen, stage, snake, food)
}

func render(screen *bufio.Writer, stage *stage, snake *snake, food *food) {
	for {
		screenClear(screen)
		if snake.checkCollidingSelf() {
			fmt.Println("Game over, Your score is ", len(snake.body)-1)
			exit()
		}
		snake.move(stage, food)
		fmt.Println(len(snake.getCoords()), snake.body, snake.getCoords())
		screenWrite(screen, stage.draw(append(snake.getCoords(), food.getCoords()...)))
		screenFlush(screen)
		time.Sleep(snake.speed)
	}
}

func exit() {
	deCleanScreen()
	os.Exit(0)
}
