package elements

import (
	"log"
	"os"

	"github.com/Matias-Barrios/SDL_Universe/SDL"
	"github.com/Matias-Barrios/SDL_Universe/board"
	"github.com/Matias-Barrios/SDL_Universe/definitions"
	"github.com/Matias-Barrios/SDL_Universe/pieces"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var FONTS map[string]*ttf.Font

func NextPieceBox(r *sdl.Renderer, p pieces.Piece) {
	var x = float64(board.Board.X + 10 + (definitions.Screen.BlockSizeW * 10))
	var y = float64(board.Board.X + (definitions.Screen.BlockSizeH * 2))
	var width = float64(definitions.Screen.Width) * 0.30
	var height = float64(definitions.Screen.Height) * 0.30
	SDL.DrawStuff(r,
		SDL.Messages_Textures["frame"],
		definitions.PointsToRatioH(x),
		definitions.PointsToRatioV(y),
		definitions.PointsToRatioV(width),
		definitions.PointsToRatioV(height))
	var pWidth = float64(definitions.Screen.BlockSizeW) * 0.50
	var pHeight = float64(definitions.Screen.BlockSizeH) * 0.50
	for ix, row := range p.Shape[p.Spin] {
		for sub_ix, val := range row {
			if val != 0 {
				SDL.DrawStuff(r,
					SDL.Block_Textures[SDL.Translate(val)],
					int(sub_ix*int(pWidth))+int(x)+int(width*0.45),
					int(ix*int(pHeight))+int(y)+int(height*0.25),
					int(pWidth),
					int(pHeight))
			}
		}
	}
}

func pointsBar(r *sdl.Renderer) {
	SDL.DrawStuff(r,
		SDL.Messages_Textures["points_bar"],
		int(sub_ix*int(pWidth))+int(x)+int(width*0.45),
		int(ix*int(pHeight))+int(y)+int(height*0.25),
		int(pWidth),
		int(pHeight))
}

func DrawText(text string, font *ttf.Font, r *sdl.Renderer, color sdl.Color, x int32, y int32, width int32, height int32) {
	solid, err := font.RenderUTF8Solid(text, color)
	if err != nil {
		log.Fatalln("Render Solid - ", err.Error())
	}
	defer solid.Free()
	t, err := r.CreateTextureFromSurface(solid)
	if err != nil {
		log.Fatalf("Failed to create texture: %s\n", err)
		os.Exit(5)
	}
	r.Copy(t, nil, &sdl.Rect{x, y, width, height})
}

func LoadFont(path string) *ttf.Font {
	if font, err := ttf.OpenFont(path, 32); err == nil {
		return font
	} else {
		log.Fatalln(err.Error())
	}
	return nil // Unreachable
}

func init() {
	if err := ttf.Init(); err != nil {
		log.Fatalf("TTF init : %s\n", err.Error())
	}
	FONTS = make(map[string]*ttf.Font)
	FONTS["normal"] = LoadFont("fonts/test.ttf")
	FONTS["8bitw"] = LoadFont("fonts/8bitw.ttf")

}
