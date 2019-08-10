package elements

import (
	"github.com/Matias-Barrios/SDL_Universe/SDL"
	"github.com/Matias-Barrios/SDL_Universe/board"
	"github.com/Matias-Barrios/SDL_Universe/definitions"
	"github.com/Matias-Barrios/SDL_Universe/pieces"
	"github.com/veandco/go-sdl2/sdl"
)

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
