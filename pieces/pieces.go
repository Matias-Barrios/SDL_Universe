package pieces

import (
	"github.com/Matias-Barrios/SDL_Universe/SDL"
	"github.com/Matias-Barrios/SDL_Universe/board"
	"github.com/Matias-Barrios/SDL_Universe/definitions"
	"github.com/veandco/go-sdl2/sdl"
)

type Piece struct {
	PosX  int
	PosY  int
	Spin  byte
	Shape map[byte][8][8]byte
}

var Pieces = map[string]Piece{
	"linea": {
		PosX: 0,
		PosY: 0,
		Spin: 0,
		Shape: map[byte][8][8]byte{
			0: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			1: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			2: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			3: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 1, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	},
}

func (p *Piece) Draw(r *sdl.Renderer, t *sdl.Texture) {
	for ix, row := range p.Shape[p.Spin] {
		for sub_ix, val := range row {
			if val != 0 {
				SDL.DrawStuff(r, t, (int32(sub_ix)*definitions.Screen.BlockSize)+(int32(p.PosX)*definitions.Screen.BlockSize)+board.Board.X, (int32(ix)*definitions.Screen.BlockSize)+(int32(p.PosY)*definitions.Screen.BlockSize)+board.Board.Y, definitions.Screen.BlockSize, definitions.Screen.BlockSize)
			}
		}
	}

}

func (p *Piece) Fall() {
	var potential_piece = *p
	potential_piece.PosY += 1
	if Fits(potential_piece) {
		p.PosY += 1
	} else {
		*p = Pieces["linea"]
		p.Spin = 1
	}

}

func Fits(p Piece) bool {
	for ix, row := range p.Shape[p.Spin] {
		for sub_ix, val := range row {
			if ix+p.PosY < len(board.Board.Cells) && sub_ix+p.PosX < len(board.Board.Cells[0]) && ix+p.PosY > -1 && sub_ix+p.PosX > -1 {
				if val != 0 && board.Board.Cells[ix+p.PosY][sub_ix+p.PosX] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func Fuse(p *Piece) {
	for ix, row := range p.Shape[p.Spin] {
		for sub_ix, val := range row {
			if ix+p.PosY < len(board.Board.Cells) && sub_ix+p.PosX < len(board.Board.Cells[0]) && ix+p.PosY > -1 && sub_ix+p.PosX > -1 {
				if val != 0 {
					board.Board.Cells[ix+p.PosY][sub_ix+p.PosX] = val
				}
			}
		}
	}
}
