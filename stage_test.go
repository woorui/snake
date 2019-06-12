package main

import (
	"fmt"
	"reflect"
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

func byteSliceEqual(a, b []byte) bool {
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_newStage_draw(t *testing.T) {
	width1, height1 := 3, 3
	stage1 := newStage(width1, height1)

	if !byteSliceEqual(stage1.matrix, []byte{124, 45, 124, 10, 124, 32, 124, 10, 124, 45, 124, 10}) ||
		!reflect.DeepEqual(stage1.mapping, map[float64]int{4: 8, 7: 9, 0: 0, 1: 1, 3: 2, 1.5: 4, 3.5: 5, 6.5: 6, 11: 10}) {
		t.Errorf("newStage(%d, %d) return mapping is %v, matrix is %v. not match except", width1, height1, stage1.mapping, stage1.matrix)
	}

	coords1 := []coordinate{
		{x: 0, y: 0},
		{x: 1, y: 2},
		{x: 5, y: 5},
	}

	m1 := stage1.draw(coords1)
	except1 := []byte{0, 45, 124, 10, 124, 32, 124, 10, 124, 0, 124, 10}
	if !byteSliceEqual(m1, except1) {
		t.Errorf("newStage(%d, %d).draw(%v) return %v except %v", width1, height1, coords1, m1, except1)
	}
	fmt.Println(m1)
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
