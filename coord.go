package main

import (
	"math/rand"
	"time"
)

// Coord is coord
type Coord struct {
	ink byte
	x   int
	y   int
}

// ICoordCollection is a coord collection
type ICoordCollection interface {
	getCoordList() CoordList
}

// CoordList is an array of Coord
type CoordList []Coord

func init() {
	rand.Seed(time.Now().UnixNano())
}

// NewCoord return a new coord within NewCoord minX, maxX, minY, maxY.
// If new coord within ban area, generate it again.
func NewCoord(minX, maxX, minY, maxY int, ink byte, ban []Coord) Coord {
	x := randRange(minX, maxX)
	y := randRange(minY, maxY)

	for _, v := range ban {
		if v.x == x && v.y == y {
			return NewCoord(minX, maxX, minY, maxY, ink, ban)
		}
	}

	return Coord{x: x, y: y, ink: ink}
}

func (list *CoordList) shift() Coord {
	if len(*list) == 0 {
		return Coord{}
	}
	s := []Coord(*list)[0]
	*list = []Coord(*list)[1:]
	return s
}

func (list *CoordList) push(ele Coord) {
	*list = append(*list, ele)
}

func (list *CoordList) contain(ele Coord) bool {
	if len(*list) == 0 {
		return false
	}
	for _, item := range *list {
		if ele.x == item.x && ele.y == item.y {
			return true
		}
	}
	return false
}

func (list CoordList) concat(others CoordList) CoordList {
	return append(list, others...)
}

// randRange returns an int >= min, < max
func randRange(min, max int) int {
	return min + rand.Intn(max-min)
}
