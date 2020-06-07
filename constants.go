package main

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

const defaultGameHeight = 25

const defaultGameWidth = 12
