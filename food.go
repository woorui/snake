package main

var (
	charFood = byte('o')
)

// Food is a random point within stage
type Food struct {
	coord Coord
}

// NewFood generate a random point within a scope
// args ban is that food can't locate in
func NewFood(minX, maxX, minY, maxY int, ban []Coord) *Food {
	c := NewCoord(minX, maxX, minY, maxY, charFood, ban)
	return &Food{coord: c}
}

func (food *Food) newLocate(minX, maxX, minY, maxY int, ban []Coord) {
	c := NewCoord(minX, maxX, minY, maxY, charFood, ban)

	food.coord.x = c.x
	food.coord.y = c.y
}

func (food *Food) getCoordList() CoordList {
	var coords CoordList
	coords.push(food.coord)
	return coords
}
