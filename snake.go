package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	charSnakeBody = byte('*')
)

// direction is the direction of snake
type direction int8

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
	mu        sync.Mutex
	head      coordinate
	body      []coordinate
	gradient  time.Duration
	ticker    *time.Ticker
	rateLimit chan struct{}
	direct    direction
	used      int
}

func newSnake() *snake {
	initialSpeed := 120 * time.Millisecond
	c := coordinate{
		ink: charSnakeBody,
		x:   2,
		y:   2,
	}
	return &snake{
		head:      c,
		body:      []coordinate{},
		rateLimit: make(chan struct{}, 1),
		ticker:    time.NewTicker(initialSpeed),
		direct:    none,
		used:      0,
	}
}

// move function make snake move.
// The snake first appeared in the top-left corner.
func (s *snake) move(stage *stage, food *food) {
	if food.coordinate.x == s.head.x && food.coordinate.y == s.head.y {
		food.newLocate(1, stage.width-1, 1, stage.height-1, append(s.getCoords(), food.getCoords()...))
		// just add one element(snake'head) to body head(not snake head)
		s.body = coordPush(s.body, s.head)
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	// change direction
	if s.direct == none {
		return
	} else if s.direct == up {
		if s.head.y == 0 {
			s.head.y = stage.height - 1
		} else {
			s.head.y--
		}
	} else if s.direct == down {
		if s.head.y == stage.height-1 {
			s.head.y = 0
		} else {
			s.head.y++
		}
	} else if s.direct == left {
		if s.head.x == 0 {
			s.head.x = stage.width - 1
		} else {
			s.head.x--
		}
	} else if s.direct == right {
		if s.head.x == stage.width-1 {
			s.head.x = 0
		} else {
			s.head.x++
		}
	}
	// remove one element from body tail, add one element(snake'head) to body head(not snake head)
	s.body = coordPush(coordShift(s.body), s.head)
}

func (s *snake) turning(direct direction) {
	s.mu.Lock()
	defer s.mu.Unlock()
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
	select {
	case <-s.rateLimit:
		switch <-input {
		case 119:
			s.turning(up)
			break
		case 115:
			s.turning(down)
			break
		case 97:
			s.turning(left)
			break
		case 100:
			s.turning(right)
			break
		}
	default:
		fmt.Println("----")
	}
}

func (s *snake) getCoords() []coordinate {
	coords := []coordinate{s.head}
	return append(coords, s.body...)
}
