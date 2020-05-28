package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	watchInterrupt(exit)

	screen := bufio.NewWriter(os.Stdout)
	stage := newStage(50, 25)
	snake := newSnake()
	// Can not locate in stage's border
	food := newFood(1, stage.width-1, 1, stage.height-1, snake.getCoords())

	nonOutputAndNobuffer()

	go func() {
		snake.rateLimit <- struct{}{}
	}()

	go snake.adapt(watchInput())

	render(screen, stage, snake, food)
}

func render(screen *bufio.Writer, stage *stage, snake *snake, food *food) {
	for range snake.ticker.C {
		go func() {
			snake.rateLimit <- struct{}{}
		}()
		fmt.Println("=====0")
		// screenClear(screen)
		if snake.checkCollidingSelf() || len(snake.body) >= (stage.width-2)*(stage.height-2) {
			fmt.Println("Game over, Your score is ", len(snake.body)-1)
			exit()
		}
		snake.move(stage, food)
		screenWrite(screen, stage.draw(append(snake.getCoords(), food.getCoords()...)))
		screenFlush(screen)
	}
}
