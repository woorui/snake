package main

import (
	"errors"
	"time"
)

// CharSnakeBody is used to render snake's body.
const CharSnakeBody = byte('*')

// CharFood is used to render food.
const CharFood = byte('o')

// CharTopBottomStageBorder is used to render top and bottom stage's border.
const CharTopBottomStageBorder = byte('-')

// CharLeftRightStageBorder is used to render left and right stage's border.
const CharLeftRightStageBorder = byte('|')

// CharBlank is used to render stage's blank.
const CharBlank = byte(' ')

// CharBreaker is used to render top and bottom stage's breaker.
const CharBreaker = byte('\n')

// CharClear is used to clear screen when print it.
var CharClear = []byte("\033[2J")

const defaultGameHeight = 12

const defaultGameWidth = 25

// ErrPlatformDontSupport said that platform don't support snake
var ErrPlatformDontSupport = errors.New("your platform don't support snake")

var defaultSpeed = 220 * time.Millisecond
