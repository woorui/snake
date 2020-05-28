package main

var (
	charFood = byte('o')
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
