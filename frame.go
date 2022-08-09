package main

import (
	"bufio"
	"io"
)

// CharClear, if you write CharClean to os.stdout, the terminal will be cleaned.
var CharClear = []byte("\033[2J")

type Framer interface {
	RenderWithCoordinate(items ...CoordinateLister) error
	Clear() error
}

// Frame is where snake and food performs.
type Frame struct {
	w         *bufio.Writer
	width     int
	height    int
	matrix    []byte
	cantorMap map[float64]int // mapping matrix index, you can get matrix item quickly.
}

// NewFrame return a stage.
func NewFramer(w io.Writer, width, height int) Framer {
	return &Frame{
		w:         bufio.NewWriter(w),
		width:     width,
		height:    height,
		matrix:    make([]byte, (width+1)*height),
		cantorMap: genCantorMap(width, height),
	}
}

func genCantorMap(width, height int) map[float64]int {
	var (
		size      = (width + 1) * height
		cantorMap = make(map[float64]int, size)
	)

	idx := 0
	for y := 0; y < height; y++ {
		for x := 0; x <= width; x++ {
			cantorMap[cantorPairingFn(x, y)] = idx
			idx++
		}
	}

	return cantorMap
}

func (f *Frame) resetMatrix() []byte {
	var (
		width  = f.width
		height = f.height
	)
	idx := 0

	for y := 0; y < height; y++ {
		for x := 0; x <= width; x++ {

			point := CharBackground

			if y == height-1 {
				point = CharBottomBorder
			} else if y == 0 {
				point = CharTopBorder
			} else if x == 0 {
				point = CharLeftBorder
			} else if x == width-1 {
				point = CharRightBorder
			}
			if x == width {
				point = CharNewLine
			}

			f.matrix[idx] = point
			idx++
		}
	}

	return f.matrix
}

func (f *Frame) RenderWithCoordinate(items ...CoordinateLister) error {
	matrix := f.resetMatrix()

	for _, item := range items {
		for _, c := range item.CoordinateList() {
			cantor := cantorPairingFn(c.x, c.y)
			idx := f.cantorMap[cantor]
			matrix[idx] = c.ink
		}
	}

	_, err := f.w.Write(matrix)
	if err != nil {
		return err
	}

	return f.w.Flush()
}

func (f *Frame) Clear() error {
	if _, err := f.w.Write(CharClear); err != nil {
		return err
	}
	return f.w.Flush()
}

// cantorPairing generator an unique number with two number.
// It works as hash function,
// More info to see: https://en.wikipedia.org/wiki/Pairing_function#Cantor_pairing_function
func cantorPairingFn(a, b int) float64 { return float64((a+b)*(a+b+1)+b) / 2 }
