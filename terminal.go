package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
)

var (
	charClear     = []byte{27, 91, 50, 74} // byte from string "\033[2J"
	charLineBreak = byte(10)               // byte from string "\n"
)

// ScreenClear make terminal screen clear, The effect is the same as command "clear"
func ScreenClear(screen *bufio.Writer) (int, error) {
	return screen.Write(charClear)
}

// ScreenWrite write content to buffer
func ScreenWrite(screen *bufio.Writer, b []byte) (int, error) {
	return screen.Write(append(b, charLineBreak))
}

// ScreenFlush write content from buffer to terminal screen
func ScreenFlush(screen *bufio.Writer) error {
	return screen.Flush()
}

// watchInterrupt watch program is interrupted and exec fn method
func watchInterrupt(fn func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println(sig)
			fn()
		}
	}()
}

func keyPress(input chan byte) {
	cleanScreen()
	b := make([]byte, 1)
	for {
		os.Stdin.Read(b)
		input <- b[0]
	}
}

// cleanScreen make terminal not show key press input
func cleanScreen() {
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
}

// deCleanScreen make terminal show key press input It is the counter operation to cleanScreen
func deCleanScreen() {
	exec.Command("stty", "-F", "/dev/tty", "-cbreak", "min", "1").Run()
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
}
