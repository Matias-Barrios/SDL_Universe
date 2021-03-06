package definitions

// Screen deinitions

const (
	screenWidth  = 1200
	screenHeight = 900
	ratioH       = float64(screenWidth) / 1000.0
	ratioV       = float64(screenHeight) / 1000.0
)

type screen struct {
	FPS        float64
	Width      int
	Height     int
	BlockSizeW int
	BlockSizeH int
}

var Screen screen

type game struct {
	Gravity float64
	Running bool
	Points  int
	Lines   int
	Goal    int
	Level   int
}

var Game = &game{
	Gravity: 5,
	Running: true,
	Points:  0,
	Lines:   0,
	Goal:    20,
	Level:   1,
}

func PointsToRatioH(h float64) int {
	return int(float64(h) * ratioH)
}

func PointsToRatioV(v float64) int {
	return int(float64(v) * ratioV)
}

func init() {
	Screen = screen{
		FPS:        1.0 / 60.0 * 1000.0,
		Width:      screenWidth,
		Height:     screenHeight,
		BlockSizeW: 40,
		BlockSizeH: 40,
	}
}
