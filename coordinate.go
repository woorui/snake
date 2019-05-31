package main

import (
	"math/rand"
	"time"
)

type coordinate struct {
	ink []byte
	x   int
	y   int
}

func coordShift(list []coordinate) []coordinate {
	return list[1:]
}

func coordPush(list []coordinate, ele coordinate) []coordinate {
	return append(list, ele)
}

// coordContain determine whether list contains ele
func coordContain(list []coordinate, ele coordinate) bool {
	for _, item := range list {
		if ele.x == item.x && ele.y == item.y {
			return true
		}
	}
	return false
}

func randXY(maxX, maxY int) (x, y int) {
	rand.Seed(time.Now().UnixNano())
	x = rand.Intn(maxX)
	y = rand.Intn(maxY)
	return
}
