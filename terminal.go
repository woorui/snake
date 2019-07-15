package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

var (
	charClear     = []byte{27, 91, 50, 74} // byte from string "\033[2J"
	charLineBreak = byte(10)               // byte from string "\n"
)

// screenClear make terminal screen clear, The effect is the same as command "clear"
func screenClear(screen *bufio.Writer) (int, error) {
	return screen.Write(charClear)
}

// screenWrite write content to buffer
func screenWrite(screen *bufio.Writer, b []byte) (int, error) {
	return screen.Write(append(b, charLineBreak))
}

// screenFlush write content from buffer to terminal screen
func screenFlush(screen *bufio.Writer) error {
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

func keyPress(input chan byte, interval time.Duration) {
	cleanScreen()

	t := newThrottle(interval)
	b := make([]byte, 1)
	for {
		os.Stdin.Read(b)
		t.exec(func() {
			input <- b[0]
		})
	}
}

// cleanScreen make terminal no buffering and no visible output
func cleanScreen() error {
	err := exec.Command("/bin/stty", "-f", "/dev/tty", "cbreak", "min", "1").Run()
	err = exec.Command("/bin/stty", "-f", "/dev/tty", "-echo").Run()

	return err
}

// deCleanScreen is the counter operation to cleanScreen
func deCleanScreen() error {
	err := exec.Command("/bin/stty", "-f", "/dev/tty", "-cbreak", "min", "1").Run()
	err = exec.Command("/bin/stty", "-f", "/dev/tty", "echo").Run()

	return err
}
