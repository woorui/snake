package main

import (
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var ErrTooManyRandTimes = errors.New("too many random times")

// Coordinate.
type Coordinate struct {
	ink byte
	x   int
	y   int
}

// NewCoordinate return a new coordinate within minX, maxX, minY, maxY,
// the new coordinate will be rendered by ink,
// If the new coordinate locates forbid, redo the NewCoordinate.
func NewCoordinate(
	minX, maxX int,
	minY, maxY int,
	ink byte,
	forbid []Coordinate) (Coordinate, error) {

	deep := 0

	return NewCoordinateWithDeep(minX, maxX, minY, maxY, ink, forbid, &deep)
}

func NewCoordinateWithDeep(minX, maxX, minY, maxY int, ink byte, forbid []Coordinate, deep *int) (Coordinate, error) {
	var (
		x = RandomInt(minX, maxX)
		y = RandomInt(minY, maxY)
	)

	*deep++

	if *deep >= 1000 {
		return Coordinate{}, ErrTooManyRandTimes
	}

	for _, v := range forbid {
		if v.x == x && v.y == y {
			return NewCoordinateWithDeep(minX, maxX, minY, maxY, ink, forbid, deep)
		}
	}

	return Coordinate{x: x, y: y, ink: ink}, nil
}

func (c Coordinate) CoordinateList() []Coordinate { return []Coordinate{c} }

// CoordinateLister
type CoordinateLister interface{ CoordinateList() []Coordinate }

func RemoveFront(list []Coordinate) []Coordinate {
	if len(list) == 0 {
		return list
	}

	return list[1:]
}

func Append(list []Coordinate, others ...Coordinate) []Coordinate { return append(list, others...) }

func Includes(list []Coordinate, element Coordinate) bool {
	for _, item := range list {
		if element.x == item.x && element.y == item.y {
			return true
		}
	}

	return false
}

// RandomInt returns an int >= min and < max.
// if min > max, return 0, returning equal vale if equals.
func RandomInt(min, max int) int {
	if min == max {
		return max
	}
	if min > max {
		return 0
	}
	return min + rand.Intn(max-min)
}
