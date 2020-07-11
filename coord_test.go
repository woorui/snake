package main

import (
	"fmt"
	"testing"
)

func genMockCoordList() CoordList {
	return CoordList([]Coord{
		{x: 1, y: 1},
		{x: 2, y: 2},
		{x: 3, y: 3},
		{x: 4, y: 4},
	})
}

var mockCoord1 = Coord{x: 4, y: 4}

var mockCoord2 = Coord{x: 5, y: 5}

func Test_coord_shift(t *testing.T) {
	coordList := genMockCoordList()
	coord := coordList.shift()
	if coord.x != 1 || coord.y != 1 {
		t.Error("coordList.shift return error")
	}
	if (coordList[0].x != 2 && coordList[0].y != 2) || len(coordList) != 3 {
		t.Error("coordList.shift test error")
	}
}

func Test_coord_push(t *testing.T) {
	coordList := genMockCoordList()
	coordList.push(mockCoord2)
	fmt.Println(coordList)
	if (coordList[len(coordList)-1].x != 5 && coordList[len(coordList)-1].y != 5) || len(coordList) != 5 {
		t.Error("coordList.push test error")
	}
}

func Test_coord_contain(t *testing.T) {
	coordList := genMockCoordList()
	c1 := coordList.contain(mockCoord1)
	c2 := coordList.contain(mockCoord2)

	if c1 != true || c2 != false {
		t.Error("coordContain test error")
	}
}
