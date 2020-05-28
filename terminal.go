package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"time"

	"errors"
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

func watchInput(input chan byte, interval time.Duration) {
	goos := runtime.GOOS
	dashF, err := polyfillDashF(goos)
	if err != nil {
		log.Fatalln("dashF: ", err)
	}
	// no buffering
	err = exec.Command("/bin/stty", dashF, "/dev/tty", "cbreak", "min", "1").Run()
	if err != nil {
		log.Fatalln("Your platform don't support snake")
	}
	// no visible output
	err = exec.Command("/bin/stty", dashF, "/dev/tty", "-echo").Run()
	if err != nil {
		log.Fatalln("Your platform don't support snake")
	}

	t := newThrottle(interval * 2)
	b := make([]byte, 1)
	for {
		os.Stdin.Read(b)
		t.exec(func() {
			input <- b[0]
		})
	}
}

// exit call when snake dead or Interrupt
func exit() {
	goos := runtime.GOOS
	dashF, err := polyfillDashF(goos)
	if err != nil {
		log.Fatalln("dashF: ", err)
	}
	err = exec.Command("/bin/stty", dashF, "/dev/tty", "-cbreak", "min", "1").Run()
	if err != nil {
		log.Fatalln("Your platform dont't support snake")
	}
	err = exec.Command("/bin/stty", dashF, "/dev/tty", "echo").Run()
	if err != nil {
		log.Fatalln("Your platform dont't support snake")
	}
	os.Exit(0)
}

func polyfillDashF(goos string) (string, error) {
	switch goos {
	case "darwin":
		return "-f", nil
	case "linux":
		return "-F", nil
	default:
		return "", errors.New("Your platform don't support snake")
	}
}
