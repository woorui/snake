package main

// Stage is where snake and food locate
type Stage struct {
	width   int
	height  int
	matrix  []byte
	mapping map[float64]int // map coord to matrix index
}

// NewStage return a stage
func NewStage(width, height int) *Stage {
	if width < 3 || height < 3 {
		width, height = defaultGameWidth, defaultGameHeight
	}
	mapping := make(map[float64]int)
	var matrix []byte
	index := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y == height-1 || y == 0 {
				matrix = append(matrix, CharTopBottomStageBorder)
			} else if x == 0 || x == width-1 {
				matrix = append(matrix, CharLeftRightStageBorder)
			} else {
				matrix = append(matrix, CharBlank)
			}
			mapping[cantorPairingFn(x, y)] = index
			index++
			if x == width-1 {
				matrix = append(matrix, charLineBreak)
				index++
			}
		}
	}
	return &Stage{
		width:   width,
		height:  height,
		matrix:  matrix,
		mapping: mapping,
	}
}

// cantorPairing generator an unique number with two number.
// It works as hash function
// More info to see: https://en.wikipedia.org/wiki/Pairing_function#Cantor_pairing_function
func cantorPairingFn(a, b int) float64 {
	num := (a+b)*(a+b+1) + b
	return float64(num) / 2
}
