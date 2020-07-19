package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
)

var (
	charClear     = []byte("\033[2J") // byte from string "\033[2J"
	charLineBreak = byte('\n')        // byte from string "\n"
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

func watchInput() chan byte {
	input := make(chan byte)
	go func() {
		b := make([]byte, 1)
		for {
			os.Stdin.Read(b)
			input <- b[0]
		}
	}()
	return input
}

func keyPressEvent() (chan os.Signal, chan Direction) {
	sig := make(chan os.Signal, 1)
	go func() {
		signal.Notify(sig, os.Interrupt)
	}()

	directionChan := make(chan Direction)
	go func() {
		bytes := make([]byte, 1)
		for {
			os.Stdin.Read(bytes)
			b := bytes[0]
			switch b {
			case 119:
				directionChan <- Up
				continue
			case 115:
				directionChan <- Down
				continue
			case 97:
				directionChan <- Left
				continue
			case 100:
				directionChan <- Right
				continue
			}
		}
	}()

	return sig, directionChan
}

// exit call when snake dead or Interrupt
func exit() {
	recoverNonOutputAndNobuffer()
	os.Exit(0)
}

func polyfillDashF(goos string) (string, error) {
	switch goos {
	case "darwin":
		return "-f", nil
	case "linux":
		return "-F", nil
	default:
		return "", ErrPlatformDontSupport
	}
}

func nonOutputAndNobuffer() error {
	goos := runtime.GOOS
	dashF, err := polyfillDashF(goos)
	if err != nil {
		return ErrPlatformDontSupport
	}
	// no buffering
	err = exec.Command("/bin/stty", dashF, "/dev/tty", "cbreak", "min", "1").Run()
	if err != nil {
		return ErrPlatformDontSupport
	}
	// no visible output
	err = exec.Command("/bin/stty", dashF, "/dev/tty", "-echo").Run()
	if err != nil {
		return ErrPlatformDontSupport
	}
	return nil
}

func recoverNonOutputAndNobuffer() error {
	goos := runtime.GOOS
	dashF, err := polyfillDashF(goos)
	if err != nil {
		return ErrPlatformDontSupport
	}
	err = exec.Command("/bin/stty", dashF, "/dev/tty", "-cbreak", "min", "1").Run()
	if err != nil {
		return ErrPlatformDontSupport
	}
	err = exec.Command("/bin/stty", dashF, "/dev/tty", "echo").Run()
	if err != nil {
		return ErrPlatformDontSupport
	}
	return nil
}
