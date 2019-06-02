package main

import "testing"

func contain(list []int, ele int) bool {
	for _, v := range list {
		if v == ele {
			return false
		}
	}
	return true
}

func Test_cantorPairing(t *testing.T) {
	tables := []struct {
		a, b, c int
	}{
		{1, 2, 7},
		{4, 5, 47},
		{3, 2, 16},
	}

	for _, table := range tables {
		res := cantorPairing(table.a, table.b)
		if res != table.c {
			t.Errorf("cantorPairing(%d,%d) was incorrect, got:%d, want:%d", table.a, table.b, res, table.c)
		}
	}

	var list []int
	for i := 0; i < 500; i++ {
		for j := 0; j < 500; j++ {
			num := cantorPairing(i, j)
			if len(list) != 0 && contain(list, num) {
				t.Error("cantorPairing func exec error")
			}
			t.Logf("i = %d, j = %d, num = %d len(list) = %d", i, j, num, len(list))
		}
	}
}
