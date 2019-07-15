package main

import (
	"time"
)

// throttle used to ensure that time-consuming tasks fire one time within a period of time
type throttle struct {
	state    bool
	returns  bool
	interval time.Duration
}

// newThrottle construct a throttle instance
func newThrottle(interval time.Duration) *throttle {
	d := &throttle{
		state:    false,
		returns:  false,
		interval: interval,
	}

	ticker := time.NewTicker(d.interval)

	go func() {
		for range ticker.C {
			d.state = true
		}
	}()

	return d
}

func (t *throttle) exec(task func()) {
	t.returns = t.state
	t.state = false

	if t.returns {
		task()
	}
}
