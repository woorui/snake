package main

import (
	"flag"
	"os"
)

var (
	Width  int
	Height int
	Speed  int
)

func init() {
	flag.IntVar(&Width, "width", 25, "game stage width")
	flag.IntVar(&Height, "height", 12, "game stage height")
	flag.IntVar(&Speed, "speed", 120, "game speed, duration between two frames")
}

func main() {
	flag.Parse()

	NewGame(os.Stdout, &GameOption{Width: Width, Height: Height, Speed: Speed}).Run()
}
