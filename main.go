package main

import (
	"time"
)

func main() {

}

func broadcastTicker(d time.Duration) (chan bool, chan struct{}) {
	directionLock := make(chan bool)
	moving := make(chan struct{})
	ticker := time.NewTicker(d)
	go func() {
		for range ticker.C {
			directionLock <- true
			moving <- struct{}{}
		}
	}()
	return directionLock, moving
}
