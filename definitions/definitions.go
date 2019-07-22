package definitions

import (
	"math"
)

// Screen deinitions

const (
	screenWidth  = 1200
	screenHeight = 900
)

type screen struct {
	Width     int
	Height    int
	BlockSize int
}

var Screen = screen{
	Width:     screenWidth,
	Height:    screenHeight,
	BlockSize: int(math.Round(3 * screenWidth / 100)),
}

type game struct {
	Gravity int
	Running bool
}

var Game = &game{
	Gravity: 10,
	Running: true,
}
