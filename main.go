package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	Version string
	Width   int
	Height  int
	Speed   int
)

func init() {
	flag.IntVar(&Width, "width", 25, "game stage width")
	flag.IntVar(&Height, "height", 12, "game stage height")
	flag.IntVar(&Speed, "speed", 120, "game speed, duration between two frames")
}

func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v" || args[1] == "version") {
		fmt.Printf("snake version: %s\n", Version)
		return
	}

	flag.Parse()
	// go run -race $(ls *.go | grep -v _test.go)
	NewGame(os.Stdout, &GameOption{Width: Width, Height: Height, Speed: Speed}).Run()
}
