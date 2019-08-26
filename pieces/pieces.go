package pieces

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/Matias-Barrios/SDL_Universe/SDL"
	"github.com/Matias-Barrios/SDL_Universe/board"
	"github.com/Matias-Barrios/SDL_Universe/definitions"
	"github.com/veandco/go-sdl2/sdl"
)

// How many milliseconds can you drift a piece before it fuses with the board ( depends on main loop delay which should be 1 millisecond)
var driftlimit int = 30

type Piece struct {
	PosX     float64
	PosY     float64
	Spin     int
	vely     int
	Drifting int
	Shape    map[int][8][8]byte
}

var AllPosiblePieces []string
var QtyOfPieces int

var Pieces = map[string]Piece{
	"line": {
		PosX:     0,
		PosY:     0,
		Spin:     0,
		Drifting: 0,
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
		PosX:     0,
		PosY:     0,
		Spin:     0,
		Drifting: 0,
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
		PosX:     0,
		PosY:     0,
		Spin:     0,
		Drifting: 0,
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
		PosX:     0,
		PosY:     0,
		Spin:     0,
		Drifting: 0,
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
					int((sub_ix*definitions.Screen.BlockSizeW)+((int(p.PosX))*definitions.Screen.BlockSizeW)+board.Board.X),
					int((ix*definitions.Screen.BlockSizeH)+(int(p.PosY))+board.Board.Y),
					int(definitions.Screen.BlockSizeW),
					int(definitions.Screen.BlockSizeH))
			}
		}
	}

}

func (p *Piece) Fall(next *Piece, c *SDL.GameContext) {
	if Fits(p, 0, 1, p.Spin) {
		p.PosY += definitions.Game.Gravity
		p.Drifting = 0
	} else {
		if p.Drifting == driftlimit {
			_, err := SDL.AUDIOS["piecedrop"].Play(-1, 0)
			if err != nil {
				log.Fatalln(err.Error())
			}
			Fuse(p, c)
			*p = *next
			*next = Pieces[RandomPiece()]
			p.Spin = 0
		} else {
			p.Drifting++
		}
	}
}

func Fits(p *Piece, velx float64, vely float64, spin int) bool {
	for ix, row := range p.Shape[spin] {
		for sub_ix, val := range row {
			if ix+(int(math.Ceil(float64(int(p.PosY)/definitions.Screen.BlockSizeH)))+int(vely)) < len(board.Board.Cells) &&
				sub_ix+(int(p.PosX)+int(velx)) < len(board.Board.Cells[0]) &&
				ix+(int(math.Ceil(float64(int(p.PosY)/definitions.Screen.BlockSizeH)))+int(vely)) > -1 && sub_ix+(int(p.PosX)+int(velx)) > -1 {
				if val != 0 && board.Board.Cells[ix+(int(math.Ceil(float64(int(p.PosY)/definitions.Screen.BlockSizeH)))+int(vely))][sub_ix+(int(p.PosX)+int(velx))] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func Fuse(p *Piece, ctx *SDL.GameContext) {
	for ix, row := range p.Shape[p.Spin] {
		for sub_ix, val := range row {
			if ix+(int(p.PosY)/definitions.Screen.BlockSizeH) < len(board.Board.Cells) && sub_ix+int(p.PosX) < len(board.Board.Cells[0]) && ix+(int(p.PosY)/definitions.Screen.BlockSizeH) > -1 && sub_ix+int(p.PosX) > -1 {
				if val != 0 {
					board.Board.Cells[ix+(int(p.PosY)/definitions.Screen.BlockSizeH)][sub_ix+int(p.PosX)] = val
				}
			}
		}
	}
	cleared := board.Board.GetFilled()
	if cleared != nil {
		if !SDL.IsPlaying(3) {
			fmt.Println("Is it Playing?")
			SDL.AUDIOS["clearedLine"].Play(3, 1)
		}
		ctx.StopMovement = true
		ctx.ClearLines = false
		for _, rowCleared := range cleared {
			ctx.ANIMATIONS = append(ctx.ANIMATIONS, &SDL.Animable{
				Posx:     board.Board.X + definitions.Screen.BlockSizeW,
				Posy:     board.Board.Y + (rowCleared * definitions.Screen.BlockSizeH),
				Width:    (10 * definitions.Screen.BlockSizeW),
				Height:   definitions.Screen.BlockSizeH,
				Textures: SDL.BeamTextures,
				Timings:  []int{10, 10, 10, 50},
				Tick:     0,
				Index:    0,
				Endless:  false,
				Finished: false,
				Handler: func() {
					ctx.StopMovement = false
					ctx.ClearLines = true
				},
			})
		}
	}
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

func (p *Piece) Move(velx float64) {
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
