package main

import (
	"time"
)

var (
	charSnakeBody = byte(42) // byte from string "*"
)

// direction is the direction of snake
type direction int

// enum direction
const (
	none  direction = 0
	up    direction = 1
	down  direction = -1
	left  direction = 2
	right direction = -2
)

// snake is the snake that can moving
type snake struct {
	head   coordinate
	body   []coordinate
	speed  time.Duration
	direct direction
}

func newSnake() *snake {
	c := coordinate{
		ink: charSnakeBody,
		x:   2,
		y:   2,
	}
	return &snake{
		head:   c,
		body:   []coordinate{c},
		speed:  time.Millisecond * 100,
		direct: none,
	}
}

// move function make snake move.
// The snake first appeared in the top-left corner.
func (s *snake) move(stage *stage, food *food) {

	if food.coordinate.x == s.head.x && food.coordinate.y == s.head.y {
		food.newLocate(stage.width, stage.height)
		// just add one element(snake'head) to body head(not snake head)
		s.body = coordPush(s.body, s.head)
	}

	// change direction
	if s.direct == none {
		return
	} else if s.direct == up {
		if s.head.y == 0 {
			s.head.y = stage.height
		} else {
			s.head.y--
		}
	} else if s.direct == down {
		if s.head.y == stage.height {
			s.head.y = 0
		} else {
			s.head.y++
		}
	} else if s.direct == left {
		if s.head.x == 0 {
			s.head.x = stage.width
		} else {
			s.head.x--
		}
	} else if s.direct == right {
		if s.head.x == stage.width {
			s.head.x = 0
		} else {
			s.head.x++
		}
	}
	// remove one element from body tail, add one element(snake'head) to body head(not snake head)
	s.body = coordPush(coordShift(s.body), s.head)
}

func (s *snake) turning(direct direction) {
	currentDirect := s.direct
	// Not change when negative or positive directiton
	if currentDirect == direct || currentDirect+direct == 0 {
		return
	}
	s.direct = direct
}

// checkCollidingSelf use to check colliding snakeSelf
func (s *snake) checkCollidingSelf() bool {
	body := s.body
	head := s.head
	bodySize := len(body)

	return bodySize >= 3 && coordContain(body[:bodySize-2], head)
}

// adapt translate input byte to snake direction, This function needs to be called asynchronously
func (s *snake) adapt(input chan byte) {
	for i := range input {
		if i == 119 {
			s.turning(up)
		} else if i == 115 {
			s.turning(down)
		} else if i == 97 {
			s.turning(left)
		} else if i == 100 {
			s.turning(right)
		}
	}
}

func (s *snake) getCoords() []coordinate {
	coords := []coordinate{s.head}
	return append(coords, s.body...)
}
