package pieces

import (
	"math/rand"
	"time"

	"github.com/Matias-Barrios/SDL_Universe/SDL"
	"github.com/Matias-Barrios/SDL_Universe/board"
	"github.com/Matias-Barrios/SDL_Universe/definitions"
	"github.com/veandco/go-sdl2/sdl"
)

type Piece struct {
	PosX  int
	PosY  int
	Spin  int
	Shape map[int][8][8]byte
}

var AllPosiblePieces []string
var QtyOfPieces int

var Pieces = map[string]Piece{
	"line": {
		PosX: 0,
		PosY: 0,
		Spin: 0,
		Shape: map[int][8][8]byte{
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
	"square": {
		PosX: 0,
		PosY: 0,
		Spin: 0,
		Shape: map[int][8][8]byte{
			0: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			1: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			2: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			3: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	},
	"L": {
		PosX: 0,
		PosY: 0,
		Spin: 0,
		Shape: map[int][8][8]byte{
			0: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			1: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			2: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			3: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 1, 0, 0},
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
	if Fits(p, 0, 1, p.Spin) {
		p.PosY += 1
	} else {
		Fuse(p)
		*p = Pieces[RandomPiece()]
		p.Spin = 0
	}

}

func Fits(p *Piece, velx int, vely int, spin int) bool {
	for ix, row := range p.Shape[spin] {
		for sub_ix, val := range row {
			if ix+(p.PosY+vely) < len(board.Board.Cells) && sub_ix+(p.PosX+velx) < len(board.Board.Cells[0]) && ix+(p.PosY+vely) > -1 && sub_ix+(p.PosX+velx) > -1 {
				if val != 0 && board.Board.Cells[ix+(p.PosY+vely)][sub_ix+(p.PosX+velx)] != 0 {
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
	board.Board.ClearLines()
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	for k, _ := range Pieces {
		AllPosiblePieces = append(AllPosiblePieces, k)
	}
	QtyOfPieces = len(AllPosiblePieces)
}

func RandomPiece() string {
	return AllPosiblePieces[rand.Intn(QtyOfPieces)]
}

func (p *Piece) Move(velx int) {
	if Fits(p, velx, 0, p.Spin) {
		p.PosX += velx
	} else {
		// TODO : Sound here
	}
}

func (p *Piece) SpinIt(spin int) {
	var next int
	if p.Spin+spin > 3 {
		next = 0
	} else if p.Spin+spin < 0 {
		next = 3
	} else {
		next = p.Spin + spin
	}
	if Fits(p, 0, 0, next) {
		p.Spin = next
	}

}
