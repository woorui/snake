package main

import (
	"sync/atomic"
)

// Direction is the direction of snake
type Direction int8

// enum direction
const (
	None  Direction = 0
	Up    Direction = 1
	Down  Direction = -1
	Left  Direction = 2
	Right Direction = -2
)

// DirectionLock lock snake direction enum
type DirectionLock int32

// enum DirectionLock
const (
	Locked   DirectionLock = 1
	UnLocked DirectionLock = 0
)

// Snake is the snake that can moving
type Snake struct {
	head          Coord
	body          CoordList
	directionChan chan Direction
	direction     Direction
	directionLock int32
}

// NewSnake return a snake
func NewSnake(x, y int, ink byte) *Snake {
	snake := &Snake{
		head:          Coord{ink, x, y},
		body:          CoordList{},
		direction:     None,
		directionLock: 0,
	}

	return snake
}

// IsBiteSelf compute whether snake eat Itself.
// if true. Game over.
func (snake *Snake) IsBiteSelf() bool {
	return len(snake.body) >= 3 && snake.body.contain(snake.head)
}

func (snake *Snake) changeDirection(direction Direction) {
	if snake.isDirectionLocked() {
		return
	}
	snake.lockDirection()

	cur := snake.direction

	if cur == direction || cur+direction == 0 {
		return
	}

	snake.direction = direction
}

// Move make snake move. need synchronous execution
func (snake *Snake) Move(maxHeight, maxWidth int) {
	switch snake.direction {
	case None:
		return
	case Up:
		if snake.head.y == 0 {
			snake.head.y = maxHeight
		} else {
			snake.head.y--
		}
	case Down:
		if snake.head.y == maxHeight {
			snake.head.y = 0
		} else {
			snake.head.y++
		}
	case Left:
		if snake.head.x == 0 {
			snake.head.x = maxWidth
		} else {
			snake.head.x--
		}
	case Right:
		if snake.head.x == maxWidth {
			snake.head.x = 0
		} else {
			snake.head.x++
		}
	}

	snake.body.shift()
	snake.body.push(snake.head)
}

func (snake *Snake) getCoords() CoordList {
	coords := []Coord{snake.head}
	return append(coords, snake.body...)
}

func (snake *Snake) lockDirection() {
	atomic.SwapInt32(&snake.directionLock, 1)
}
func (snake *Snake) unLockDirection() {
	atomic.SwapInt32(&snake.directionLock, 0)
}

func (snake *Snake) isDirectionLocked() bool {
	return snake.directionLock == 1
}
