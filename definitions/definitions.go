package definitions

import "math"

// Screen deinitions

const (
	screenWidth  = 800
	screenHeight = 600
)

type screen struct {
	Width     int32
	Height    int32
	BlockSize int32
}

var Screen = screen{
	Width:     screenWidth,
	Height:    screenHeight,
	BlockSize: int32(math.Round(3 * screenWidth / 100)),
}
