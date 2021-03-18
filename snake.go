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

// Snake is the snake that can moving
type Snake struct {
	head                Coord
	body                *CoordList
	directionController *directionController
}

type directionController struct {
	cur  Direction
	pool []Direction
}

func (dc *directionController) changeDirection(direction Direction) {
	last := None
	if len(dc.pool) != 0 {
		last = dc.pool[len(dc.pool)-1]
	}
	if last == direction || last+direction == 0 {
		return
	}
	dc.pool = append(dc.pool, direction)
}

func (dc *directionController) apply() {
	pool := dc.pool
	dc.pool = []Direction{}
	for i := range pool {
		last := pool[len(pool)-i-1]
		if last != dc.cur && last+dc.cur != 0 {
			dc.cur = last
			return
		}
	}
}

// NewSnake return a snake
func NewSnake(x, y int, ink byte) *Snake {
	snake := &Snake{
		head:                Coord{ink, x, y},
		body:                &CoordList{},
		directionController: &directionController{},
	}

	return snake
}

// IsBiteSelf compute whether snake eat Itself.
// if true. Game over.
func (snake *Snake) IsBiteSelf() bool {
	size := snake.body.size()
	if size >= 3 {
		realBody := CoordList([]Coord(*snake.body)[:snake.body.size()-2])
		return realBody.contain(snake.head)
	}
	return false
}

func (snake *Snake) apply() {
	snake.directionController.apply()
}

func (snake *Snake) changeDirection(direction Direction) {
	snake.directionController.changeDirection(direction)
}

// Move make snake move. need synchronous execution
func (snake *Snake) Move(maxHeight, maxWidth int) {
	switch snake.directionController.cur {
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
	return append(coords, *snake.body...)
}
