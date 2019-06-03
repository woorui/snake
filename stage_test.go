package main

import (
	"testing"
)

func contain(list []float64, ele float64) bool {
	for _, v := range list {
		if v == ele {
			return true
		}
	}
	return false
}

func Test_cantorPairing(t *testing.T) {
	tables := []struct {
		a, b int
		c    float64
	}{
		{0, 1, 1.5},
		{1, 0, 1},
		{3, 2, 16},
	}

	for _, table := range tables {
		res := cantorPairing(table.a, table.b)
		if res != table.c {
			t.Errorf("cantorPairing(%d,%d) was incorrect, got:%f, want:%f", table.a, table.b, res, table.c)
		}
	}

	var list []float64
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if i != j {
				if cantorPairing(i, j) == cantorPairing(j, i) {
					t.Errorf("cantorPairing func exec error cantorPairing(%d, %d) = %f cantorPairing(%d, %d) = %f", i, j, cantorPairing(i, j), i, i, cantorPairing(j, i))
				}
			}
			num := cantorPairing(i, j)
			list = append(list, num)
			if len(list) > 2 {
				if contain(list[:len(list)-2], num) {
					t.Errorf("cantorPairing func exec error list = %v, num = %fs", list, num)
					break
				}
			}
		}
	}
}
