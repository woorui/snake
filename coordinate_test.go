package main

import "testing"

func genMockCoordinateList() []coordinate {
	return []coordinate{
		coordinate{x: 1, y: 1},
		coordinate{x: 2, y: 2},
		coordinate{x: 3, y: 3},
		coordinate{x: 4, y: 4},
	}
}

var mockCoordinate1 = coordinate{x: 4, y: 4}

var mockCoordinate2 = coordinate{x: 5, y: 5}

func Test_coordShift(t *testing.T) {
	cs := coordShift(genMockCoordinateList())
	if (cs[0].x != 2 && cs[0].y != 2) || len(cs) != 3 {
		t.Error("coordShift exec error")
	}
}

func Test_coordPush(t *testing.T) {
	cs := coordPush(genMockCoordinateList(), mockCoordinate2)
	if (cs[len(cs)-1].x != 5 && cs[len(cs)-1].y != 5) || len(cs) != 5 {
		t.Error("coordPush exec error")
	}
}

func Test_coordContain(t *testing.T) {
	isContain1 := coordContain(genMockCoordinateList(), mockCoordinate1)
	isContain2 := coordContain(genMockCoordinateList(), mockCoordinate2)

	if isContain1 != true || isContain2 != false {
		t.Error("coordContain exec error")
	}
}
