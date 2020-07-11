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
	stage := newStage(25, 12)
	snake := newSnake()
	// Can not locate in stage's border
	food := newFood(1, stage.width-1, 1, stage.height-1, snake.getCoords())

	nonOutputAndNobuffer()

	go snake.turningInchannel(watchInput())

	directionLock, moving := broadcastTicker(120 * time.Millisecond)

	for {
		select {
		case <-moving:
			screenClear(screen) // race here
			if snake.checkCollidingSelf() || len(snake.body) >= (stage.width-2)*(stage.height-2) {
				fmt.Println("Game over, Your score is ", len(snake.body)-1)
				exit()
			}
			snake.move(stage, food)

			screenWrite(screen, stage.draw(append(snake.getCoords(), food.getCoords()...)))
			screenFlush(screen)
		case <-directionLock:
			snake.unLockDirection()
		default:
		}
	}
}

func broadcastTicker(d time.Duration) (chan bool, chan struct{}) {
	directionLock := make(chan bool)
	moving := make(chan struct{})
	ticker := time.NewTicker(d)
	go func() {
		for range ticker.C {
			directionLock <- true
			moving <- struct{}{}
		}
	}()
	return directionLock, moving
}
