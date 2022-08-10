package main

const (
	CharTopBorder    = byte('-')  // CharTopBorder is used to render the top border.
	CharBottomBorder = byte('-')  // CharBottomBorder is used to render the bottom border.
	CharLeftBorder   = byte('|')  // CharLeftBorder is used to render the left border.
	CharRightBorder  = byte('|')  // CharRightBorder is used to render the right border.
	CharBackground   = byte(' ')  // CharBackground is used to render the background.
	CharFood         = byte('o')  // CharFood is used to render food.
	CharNewLine      = byte('\n') // CharNewLine was written when terminal new line.
	CharSnakeBody    = byte('*')  // CharSnakeBody is used to render snake's body.
	CharSnakeHead    = byte('*')  // CharSnakeHead is used to render snake's head.
)
