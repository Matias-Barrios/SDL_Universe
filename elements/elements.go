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
}
