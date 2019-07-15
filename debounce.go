package main

import (
	"time"
)

// debounce used to ensure that time-consuming tasks do not fire so often
type debounce struct {
	state   bool
	returns bool
	delay   time.Duration
}

// newDebounce construct a debounce instance
func newDebounce(delay time.Duration) *debounce {
	d := &debounce{
		state:   false,
		returns: false,
		delay:   delay,
	}

	ticker := time.NewTicker(d.delay)

	go func() {
		for range ticker.C {
			d.state = true
		}
	}()

	return d
}

func (d *debounce) exec(task func()) {
	d.returns = d.state
	d.state = false

	if d.returns {
		task()
	}
}
