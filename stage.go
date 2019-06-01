package main

// stage is where snake and food locate
type stage struct {
	width  int
	height int
}

func newStage(width, height int) *stage {
	return &stage{
		width:  width,
		height: height,
	}
}
