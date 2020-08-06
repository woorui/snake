package main

import "flag"

func main() {
	width := flag.Int("width", defaultGameWidth, "Game stage width.")
	height := flag.Int("height", defaultGameHeight, "Game stage height.")
	debug := flag.Bool("debug", false, "debug mode.")
	flag.Parse()

	game := NewGame(GameOpts{Width: *width, Height: *height, Debug: *debug})

	game.Run()
}
