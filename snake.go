package main

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

func (d Direction) String() string {
	switch d {
	case None:
		return "None"
	case Up:
		return "Up"
	case Down:
		return "Down"
	case Left:
		return "Left"
	case Right:
		return "Right"
	}
	return "Unknown"
}

// Controller controls snake direction,
// some direction changing should be ignored in one frame.
type Controller interface {
	Get() Direction
	Change(Direction)
}

type defaultController struct {
	last     Direction
	accepted []Direction
}

func (dc *defaultController) Change(dt Direction) {
	last := dc.last
	if len(dc.accepted) != 0 {
		last = dc.accepted[len(dc.accepted)-1]
	}
	// same direction
	if last == dt {
		return
	}
	// negative direction
	if last+dt == 0 {
		if len(dc.accepted) != 2 {
			dc.accepted = append(dc.accepted, dt)
		}
		return
	}
	dc.accepted = append(dc.accepted, dt)
}

func (dc *defaultController) Get() Direction {
	last := dc.last
	if len(dc.accepted) == 0 {
		return last
	}
	var (
		size = len(dc.accepted)
		next = None
		cur  = last
	)
	cur = dc.accepted[size-1]
	if size >= 2 {
		var (
			first  = dc.accepted[size-2]
			second = dc.accepted[size-1]
		)

		// turn back or change track
		if last+second == 0 || last == second {
			cur = first
			next = second
		}
	}

	if next != None {
		dc.accepted = dc.accepted[0:1]
		dc.accepted[0] = next
	} else {
		dc.accepted = dc.accepted[:0]
	}

	if last+cur == 0 {
		cur = last
	} else {
		dc.last = cur // set last
	}

	return cur
}

// Snake is the snake that can moving
type Snake struct {
	head       Coordinate
	body       []Coordinate
	controller Controller
}

// NewSnake return a snake.
//
// x, y is the initial position.
func NewSnake(x, y int, ink byte, controller Controller) *Snake {
	snake := &Snake{
		head:       Coordinate{ink, x, y},
		body:       []Coordinate{},
		controller: controller,
	}

	return snake
}

// IsBumpSelf compute whether snake bumps itself.
// if true. There should game over.
func (snake *Snake) IsBumpSelf() bool {
	size := len(snake.body)

	// snake can not bumps its neck.
	if size < 3 {
		return false
	}
	bumpable := snake.body[:size-2]

	return Includes(bumpable, snake.head)
}

func (snake *Snake) ChangeDirection(direction Direction) { snake.controller.Change(direction) }

// Move make snake move. need synchronous execution
func (snake *Snake) Move(minX, maxX, minY, maxY int) {
	switch snake.controller.Get() {
	case None:
		return
	case Up:
		if snake.head.y == minY {
			snake.head.y = maxY
		} else {
			snake.head.y--
		}
	case Down:
		if snake.head.y == maxY {
			snake.head.y = minY
		} else {
			snake.head.y++
		}
	case Left:
		if snake.head.x == minX {
			snake.head.x = maxX
		} else {
			snake.head.x--
		}
	case Right:
		if snake.head.x == maxX {
			snake.head.x = minX
		} else {
			snake.head.x++
		}
	}

	// delete tail, move forward head.
	snake.body = RemoveFront(snake.body)
	snake.body = Append(snake.body, snake.head)
}

func (snake *Snake) CoordinateList() []Coordinate { return append(snake.body, snake.head) }
