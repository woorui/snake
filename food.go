package main

import (
	"math/rand"
	"time"
)

var (
	charFood = []byte{27} // byte from string "o"
)

// food is a random point within stage
type food struct {
	coordinate coordinate
}

// newFood generate a random point within a scope
// args restriction is that food can't locate in
func newFood(maxX, maxY int, restriction []coordinate) *food {
	x, y := randXY(maxX, maxY)
	c := coordinate{x: x, y: y, ink: charFood}

	if coordContain(restriction, c) {
		return newFood(maxX, maxY, restriction)
	}

	return &food{coordinate: c}
}

func (f *food) newLocate(maxX, maxY int) {
	x, y := randXY(maxX, maxY)
	f.coordinate.x = x
	f.coordinate.y = y
}

func (f *food) getCoords() []coordinate {
	var coords []coordinate
	return append(coords, f.coordinate)
}

func randXY(maxX, maxY int) (x, y int) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(maxX)
	y = rand.Intn(maxY)
	return
}
