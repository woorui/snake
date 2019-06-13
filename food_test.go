package main

import (
	"math/rand"
	"testing"
	"time"
)

func Test_randRange(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		x := randRange(0, 10)
		y := randRange(0, 10)
		if x < 0 || x > 10 || y < 0 || y > 10 {
			t.Error("randXY exec error")
			return
		}
	}
}
