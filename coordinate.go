package main

import (
	"math/rand"
	"time"
)

type coordinate struct {
	ink byte
	x   int
	y   int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func newCoordinate(minX, maxX, minY, maxY int, ink byte, restriction []coordinate) coordinate {
	x := randRange(minX, maxX)
	y := randRange(minY, maxY)

	for _, v := range restriction {
		if v.x == x && v.y == y {
			return newCoordinate(minX, maxX, minY, maxY, ink, restriction)
		}
	}

	return coordinate{x: x, y: y, ink: ink}
}

func coordShift(list []coordinate) []coordinate {
	if len(list) == 0 {
		return list
	}
	return list[1:]
}

func coordPush(list []coordinate, ele coordinate) []coordinate {
	return append(list, ele)
}

// coordContain determine whether list contains ele
func coordContain(list []coordinate, ele coordinate) bool {
	if len(list) == 0 {
		return false
	}
	for _, item := range list {
		if ele.x == item.x && ele.y == item.y {
			return true
		}
	}
	return false
}

// randRange returns an int >= min, < max
func randRange(min, max int) int {
	return min + rand.Intn(max-min)
}
