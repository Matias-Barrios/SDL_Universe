package definitions

import "fmt"

// Screen deinitions

const (
	screenWidth  = 800
	screenHeight = 600
	ratioH       = float64(screenWidth / 1000)
	ratioV       = float64(screenHeight / 1000)
)

type screen struct {
	Width     int
	Height    int
	BlockSize int
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

func PointsToRatioH(h int) int {
	fmt.Println("h: ", h, "ratioH : ", ratioH)
	return int(float64(h) * ratioH)
}

func PointsToRatioV(v int) int {
	return int(float64(v) * ratioV)
}

func init() {
	Screen = screen{
		Width:     screenWidth,
		Height:    screenHeight,
		BlockSize: PointsToRatioH(25),
	}
	fmt.Println("X : ", PointsToRatioH(50))
	fmt.Println("Y : ", PointsToRatioV(50))
	fmt.Println("Block Size : ", Screen.BlockSize)
}
