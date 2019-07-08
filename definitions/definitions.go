package definitions

import (
	"github.com/Matias-Barrios/SDL_Universe/SDL"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

// Screen deinitions

const (
	screenWidth  = 800
	screenHeight = 600
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
}

var Game = &game{
	Gravity: 10,
}

var Block_Textures map[string]*sdl.Texture

func BricksLoadTextures(w *sdl.Window,r *sdl.Renderer) {
	Block_Textures = make(map["greyblock"]*sdl.Texture,0,10)
	Block_Textures["greyblock"] = SDL.GetTexture(w,r,"assets/greyblock.png")
	Block_Textures["greenblock"] = SDL.GetTexture(w,r,"assets/greenblock.png")
}

func Translate(number int){
	switch number {
	case 1 : 
		return "greyblock"
	case 2 : 
		return "greenblock"
	default : 
		log.Fatalf("%s\n",fmt.Errorf("Unknown identifier for texture"))
	}
}