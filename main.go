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
	stage := newStage(40, 20)
	snake := newSnake()
	food := newFood(stage.width, stage.height, snake.getCoords())

	input := make(chan byte)
	go keyPress(input)
	go snake.adapter(input)

	render(screen, stage, snake, food)
}

func render(screen *bufio.Writer, stage *stage, snake *snake, food *food) {
	for {
		ScreenClear(screen)
		if snake.checkCollidingSelf() {
			fmt.Println("Game over, your score is ", len(snake.body)-1)
			break
		}
		snake.move(stage, food)
		ScreenWrite(screen, stage.draw(append(snake.getCoords(), food.getCoords()...)))
		ScreenFlush(screen)
		time.Sleep(snake.speed)
	}
}

func exit() {
	deCleanScreen()
	os.Exit(0)
}
