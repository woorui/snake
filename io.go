package main

import (
	"bufio"
	"errors"
	"io"
	"os/exec"
	"runtime"
)

// ErrPlatformDontSupport said that platform don't support snake
var ErrPlatformDontSupport = errors.New("your platform don't support snake")

var DirectionMapping = map[byte]Direction{
	'w': Up,
	's': Down,
	'a': Left,
	'd': Right,
}

// dashF is used for MAC and Linux compatibility.
func dashF(goos string) (string, error) {
	switch goos {
	case "darwin":
		return "-f", nil
	case "linux":
		return "-F", nil
	default:
		return "", ErrPlatformDontSupport
	}
}

// OpenGameMode lets Terminal switch to game mode and returns close function for close
// game mode.
func OpenGameMode() (func() error, error) {
	goos := runtime.GOOS

	dashF, err := dashF(goos)
	if err != nil {
		return func() error { return nil }, ErrPlatformDontSupport
	}

	err = exec.Command("/bin/stty", dashF, "/dev/tty", "cbreak", "min", "1", "-echo").Run()
	if err != nil {
		return func() error { return nil }, ErrPlatformDontSupport
	}

	return func() error {
		return exec.Command("/bin/stty", dashF, "/dev/tty", "-cbreak", "min", "1", "echo").Run()
	}, nil
}

// ReadToDirection read from r to Direction
func ReadToDirection(r io.Reader) chan Direction {
	dch := make(chan Direction)

	go func() {
		in := bufio.NewReader(r)

		for {
			b, err := in.ReadByte()
			if err == io.EOF {
				close(dch)
				break
			}
			d, ok := DirectionMapping[b]
			if ok {
				dch <- d
			}
		}
	}()

	return dch
}
