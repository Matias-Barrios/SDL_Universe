package definitions

import (
	"fmt"
	"math"
)

// Screen deinitions

const (
	screenWidth  = 1200
	screenHeight = 900
	ratioH       = float64(screenWidth) / 1000.0
	ratioV       = float64(screenHeight) / 1000.0
)

type screen struct {
	Width      int
	Height     int
	BlockSizeW int
	BlockSizeH int
}

var Screen screen

type game struct {
	Gravity int
	Running bool
}

var Game = &game{
	Gravity: 10,
	Running: true,
}

func PointsToRatioH(h float64) int {
	return int(math.Ceil(float64(h) * ratioH))
}

func PointsToRatioV(v float64) int {
	return int(math.Ceil(float64(v) * ratioV))
}

func init() {
	Screen = screen{
		Width:      screenWidth,
		Height:     screenHeight,
		BlockSizeW: PointsToRatioH(25),
		BlockSizeH: PointsToRatioV(25),
	}
	fmt.Println(PointsToRatioH(50))
}
