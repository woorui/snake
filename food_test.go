package main

import "testing"

func Test_randXY(t *testing.T) {
	for i := 0; i < 100; i++ {
		x, y := randXY(10, 10)
		if x < 0 || x > 10 || y < 0 || y > 10 {
			t.Error("randXY exec error")
			return
		}
	}
}
