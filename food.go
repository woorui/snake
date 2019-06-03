package main

import (
	"math/rand"
	"time"
)

var (
	charFood = byte(111) // byte from string "o"
)

// food is a random point within stage
type food struct {
	coordinate coordinate
}

// newFood generate a random point within a scope
// args restriction is that food can't locate in
func newFood(minX, maxX, minY, maxY int, restriction []coordinate) *food {
	c := newCoordinate(minX, maxX, minY, maxY, charFood, restriction)
	return &food{coordinate: c}
}

func (f *food) newLocate(minX, maxX, minY, maxY int, restriction []coordinate) {
	c := newCoordinate(minX, maxX, minY, maxY, charFood, restriction)

	f.coordinate.x = c.x
	f.coordinate.y = c.y
}

func (f *food) getCoords() []coordinate {
	var coords []coordinate
	return append(coords, f.coordinate)
}

func randXY(minX, maxX, minY, maxY int) (x, y int) {
	rand.Seed(time.Now().UnixNano())
	x = randRange(minX, maxX)
	y = randRange(minY, maxY)
	return
}

func randRange(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}
