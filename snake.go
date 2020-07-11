package main

import (
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
	mu       sync.Mutex
	head     coordinate
	body     []coordinate
	gradient time.Duration
	direct   direction
	used     int
}

func newSnake() *snake {
	c := coordinate{
		ink: charSnakeBody,
		x:   2,
		y:   2,
	}
	return &snake{
		head:   c,
		mu:     sync.Mutex{},
		body:   []coordinate{},
		direct: none,
		used:   0,
	}
}

// move function make snake move.
// The snake first appeared in the top-left corner.
func (s *snake) move(stage *stage, food *food) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if food.coordinate.x == s.head.x && food.coordinate.y == s.head.y {
		food.newLocate(1, stage.width-1, 1, stage.height-1, append(s.getCoords(), food.getCoords()...))
		// just add one element(snake'head) to body head(not snake head)
		s.body = coordPush(s.body, s.head)
	}

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
	if s.isDirectionLocked() {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	// Not change when negative or positive directiton
	currentDirect := s.direct
	s.lockDirection()
	if currentDirect == direct || currentDirect+direct == 0 {
		return
	}
	s.direct = direct
}

func (s *snake) turningInchannel(input chan byte) {
	for {
		select {
		case i := <-input:
			switch i {
			case 119:
				s.tryChangeDirection(up)
				break
			case 115:
				s.tryChangeDirection(down)
				break
			case 97:
				s.tryChangeDirection(left)
				break
			case 100:
				s.tryChangeDirection(right)
				break
			}
		default:
		}
	}
}

func (s *snake) tryChangeDirection(direct direction) {
	if s.isDirectionLocked() {
		return
	}
	s.lockDirection()
	currentDirect := s.direct
	if currentDirect == direct || currentDirect+direct == 0 {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.direct = direct
}

// checkCollidingSelf use to check colliding snakeSelf
func (s *snake) checkCollidingSelf() bool {
	body := s.body
	head := s.head
	bodySize := len(body)

	return bodySize >= 3 && coordContain(body[:bodySize-2], head)
}

func (s *snake) getCoords() []coordinate {
	coords := []coordinate{s.head}
	return append(coords, s.body...)
}

// 1：lock
// 0：unlock
func (s *snake) lockDirection() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.used = 1
}
func (s *snake) unLockDirection() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.used = 0
}

func (s *snake) isDirectionLocked() bool {
	return s.used == 1
}
