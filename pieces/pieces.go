package pieces

import (
	"math"
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
	VelY  int
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
				{0, 0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			1: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 2, 2, 2, 2, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			2: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 0, 2, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			3: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 2, 2, 2, 2, 0, 0},
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
				{0, 0, 0, 3, 3, 0, 0, 0},
				{0, 0, 0, 3, 3, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			1: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 3, 3, 0, 0, 0},
				{0, 0, 0, 3, 3, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			2: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 3, 3, 0, 0, 0},
				{0, 0, 0, 3, 3, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			3: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 3, 3, 0, 0, 0},
				{0, 0, 0, 3, 3, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	},
	"l": {
		PosX: 0,
		PosY: 0,
		Spin: 0,
		Shape: map[int][8][8]byte{
			0: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 4, 0, 0, 0},
				{0, 0, 0, 0, 4, 0, 0, 0},
				{0, 0, 0, 4, 4, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			1: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 4, 0, 0, 0, 0},
				{0, 0, 0, 4, 4, 4, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			2: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 4, 4, 0, 0},
				{0, 0, 0, 0, 4, 0, 0, 0},
				{0, 0, 0, 0, 4, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			3: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 4, 4, 4, 0, 0},
				{0, 0, 0, 0, 0, 4, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	},
	"z": {
		PosX: 0,
		PosY: 0,
		Spin: 0,
		Shape: map[int][8][8]byte{
			0: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 5, 0, 0, 0},
				{0, 0, 0, 5, 5, 0, 0, 0},
				{0, 0, 0, 5, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			1: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 5, 5, 0, 0, 0, 0},
				{0, 0, 0, 5, 5, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			2: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 5, 0, 0, 0},
				{0, 0, 0, 5, 5, 0, 0, 0},
				{0, 0, 0, 5, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
			3: {
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 5, 5, 0, 0, 0, 0},
				{0, 0, 0, 5, 5, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	},
}

func (p *Piece) Draw(r *sdl.Renderer) {
	for ix, row := range p.Shape[p.Spin] {
		for sub_ix, val := range row {
			if val != 0 {
				SDL.DrawStuff(r,
					SDL.Block_Textures[SDL.Translate(val)],
					int((sub_ix*definitions.Screen.BlockSizeW)+((p.PosX)*definitions.Screen.BlockSizeW)+board.Board.X),
					int((ix*definitions.Screen.BlockSizeH)+(p.PosY)+board.Board.Y),
					int(definitions.Screen.BlockSizeW),
					int(definitions.Screen.BlockSizeH))
			}
		}
	}

}

func (p *Piece) Fall() {
	if Fits(p, 0, 1, p.Spin) {
		p.PosY += definitions.Game.Gravity
	} else {
		Fuse(p)
		*p = Pieces[RandomPiece()]
		p.Spin = 0
	}

}

func Fits(p *Piece, velx int, vely int, spin int) bool {
	for ix, row := range p.Shape[spin] {
		for sub_ix, val := range row {
			if ix+(int(math.Ceil(float64(p.PosY/definitions.Screen.BlockSizeH)))+vely) < len(board.Board.Cells) &&
				sub_ix+(p.PosX+velx) < len(board.Board.Cells[0]) &&
				ix+(int(math.Ceil(float64(p.PosY/definitions.Screen.BlockSizeH)))+vely) > -1 && sub_ix+(p.PosX+velx) > -1 {
				if val != 0 && board.Board.Cells[ix+(int(math.Ceil(float64(p.PosY/definitions.Screen.BlockSizeH)))+vely)][sub_ix+(p.PosX+velx)] != 0 {
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
			if ix+(p.PosY/definitions.Screen.BlockSizeH) < len(board.Board.Cells) && sub_ix+p.PosX < len(board.Board.Cells[0]) && ix+(p.PosY/definitions.Screen.BlockSizeH) > -1 && sub_ix+p.PosX > -1 {
				if val != 0 {
					board.Board.Cells[ix+(p.PosY/definitions.Screen.BlockSizeH)][sub_ix+p.PosX] = val
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
