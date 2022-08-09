package main

import (
	"testing"
)

func Test_Snake(t *testing.T) {
	snake := NewSnake(1, 1, 'x', &defaultController{})

	snake.ChangeDirection(Down)

	if got := snake.controller.Get(); got != Down {
		t.Errorf("got = %v, want = %v", got, Down)
	}

	// TODO: test result
	snake.Move(0, 5, 0, 5)
	snake.CoordinateList()
	snake.IsBumpSelf()
}

func Test_Controller(t *testing.T) {

	ctl := &defaultController{last: Left, accepted: []Direction{Right}}

	assertGetEqual(t, ctl, Left, "init state 1")

	assertGetEqual(t, ctl, Left, "init state 2")

	ctl.Change(Down)
	ctl.Change(Down)
	ctl.Change(Down)

	assertGetEqual(t, ctl, Down, "turn down")

	assertGetEqual(t, ctl, Down, "no changing")

	ctl.Change(Up)

	assertGetEqual(t, ctl, Down, "negtive direction")

	ctl.Change(Right)

	assertGetEqual(t, ctl, Right, "normal change")
	assertGetEqual(t, ctl, Right, "normal change")
	assertGetEqual(t, ctl, Right, "normal change")
	assertGetEqual(t, ctl, Right, "normal change")
	assertGetEqual(t, ctl, Right, "after normal change")

	ctl.Change(Left)
	assertGetEqual(t, ctl, Right, "negtive change after normal change")

	// irregular turn back
	ctl.Change(Up)
	ctl.Change(Down)
	ctl.Change(Down)
	ctl.Change(Left)
	ctl.Change(Left)

	assertGetEqual(t, ctl, Down, "irregular turn back step 1")
	assertGetEqual(t, ctl, Left, "irregular turn back step 2")
	assertGetEqual(t, ctl, Left, "irregular turn back step 3")

	ctl.Change(Down)
	ctl.Change(Left)

	assertGetEqual(t, ctl, Down, "change track step 1")
	assertGetEqual(t, ctl, Left, "change track step 2")
}

func assertGetEqual(t *testing.T, ctl Controller, d Direction, cas string) {
	if got := ctl.Get(); got != d {
		t.Errorf("case: %s got = %v, want = %v", cas, got, d)
	}
}
