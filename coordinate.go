package main

type coordinate struct {
	ink byte
	x   int
	y   int
}

func newCoordinate(minX, maxX, minY, maxY int, ink byte, restriction []coordinate) coordinate {
	x, y := randXY(minX, maxX, minY, maxY)

	if coordContain(restriction, coordinate{x: x, y: y}) {
		x, y = randXY(minX, maxX, minY, maxY)
	}

	return coordinate{ink, x, y}
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
	for _, item := range list {
		if ele.x == item.x && ele.y == item.y {
			return true
		}
	}
	return false
}
