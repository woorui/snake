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

func newCoordinate(minX, maxX, minY, maxY int, ink byte, restriction []coordinate) coordinate {
	// fmt.Println("----")
	x, y := randXY(minX, maxX, minY, maxY)

	if coordContain(restriction, coordinate{x: x, y: y}) {
		x, y = randXY(minX, maxX, minY, maxY)
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
