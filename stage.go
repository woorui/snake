package main

var (
	charTopBottomBorder = byte('-')
	charLeftRightBorder = byte('|')
	charBlank           = byte(' ')
	charBreak           = byte('\n')
)

// stage is where snake and food locate
type stage struct {
	width   int
	height  int
	matrix  []byte
	mapping map[float64]int // map coordinate to matrix index
}

func newStage(width, height int) *stage {
	if width < 3 || height < 3 {
		width, height = 50, 25
	}
	mapping := make(map[float64]int)
	var matrix []byte
	index := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if y == height-1 || y == 0 {
				matrix = append(matrix, charTopBottomBorder)
			} else if x == 0 || x == width-1 {
				matrix = append(matrix, charLeftRightBorder)
			} else {
				matrix = append(matrix, charBlank)
			}

			mapping[cantorPairing(x, y)] = index
			index++
			if x == width-1 {
				matrix = append(matrix, charBreak)
				index++
			}

		}
	}
	return &stage{
		width:   width,
		height:  height,
		matrix:  matrix,
		mapping: mapping,
	}
}

func (s *stage) draw(coords []coordinate) []byte {
	m := make([]byte, len(s.matrix))
	copy(m, s.matrix)
	for _, c := range coords {
		index := s.mapping[cantorPairing(c.x, c.y)]
		m[index] = c.ink
	}

	return m
}

// cantorPairing generator an unique number with two number, Work as hash function
// More info to see: https://en.wikipedia.org/wiki/Pairing_function#Cantor_pairing_function
func cantorPairing(a, b int) int {
	num := (a+b)*(a+b+1) + b
	return (num / 2) * 10
}
