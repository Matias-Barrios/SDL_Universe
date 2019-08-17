package board

import (
	"math/rand"

	"github.com/Matias-Barrios/SDL_Universe/SDL"
	"github.com/Matias-Barrios/SDL_Universe/definitions"
	"github.com/veandco/go-sdl2/sdl"
)

type board struct {
	X     int
	Y     int
	Cells [][]byte
}

var Board = &board{
	X: definitions.PointsToRatioH(10),
	Y: -definitions.PointsToRatioV(float64(definitions.Screen.BlockSizeH * 5)),
	Cells: [][]byte{
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 253},
		{253, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 253},
		{253, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 253},
		{253, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 253},
		{253, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 253},
		{252, 254, 254, 254, 254, 254, 254, 254, 254, 254, 254, 251},
	},
}

// this is the tall of the board minus the first row
var tall = len(Board.Cells) - 2

// This represents how many columns of a row can be cleared,
//  as the first and last are actually the left and right borders
var clearable_columns = 11

func Draw(r *sdl.Renderer) {
	for ix, row := range Board.Cells {
		for sub_ix, val := range row {
			if val != 0 && ix > 7 {
				SDL.DrawStuff(r,
					SDL.Block_Textures[SDL.Translate(val)],
					(sub_ix*definitions.Screen.BlockSizeW)+Board.X,
					(ix*definitions.Screen.BlockSizeH)+Board.Y,
					(definitions.Screen.BlockSizeW),
					(definitions.Screen.BlockSizeH))
			}
		}
	}
}

func Lose(r *sdl.Renderer, ctx *SDL.GameContext) {
	for ix, row := range Board.Cells {
		for sub_ix, val := range row {
			if ix > 7 {
				return
			} else if sub_ix < 1 || sub_ix > 10 {
				continue
			} else if val != 0 {
				ctx.StopMovement = true
				ctx.ClearLines = false
				ctx.ANIMATIONS = append(ctx.ANIMATIONS, &SDL.Animable{
					Posx:     definitions.PointsToRatioH(float64(Board.X + definitions.Screen.BlockSizeW)),
					Posy:     definitions.PointsToRatioV(float64(Board.Y + (definitions.Screen.BlockSizeH * 10))),
					Width:    definitions.PointsToRatioH(float64(definitions.Screen.BlockSizeW * 10)),
					Height:   definitions.PointsToRatioV(float64(definitions.Screen.BlockSizeH * 10)),
					Textures: SDL.YouLoseTextures,
					Timings:  []int{100, 100, 100, 100},
					Tick:     0,
					Index:    0,
					Endless:  true,
					Finished: false,
					Handler: func() {

					},
				})
			}
		}
	}

}

func (b *board) ClearLines() []int {
	var res = Board.GetFilled()
	for i := tall; i >= 0; i-- {
		if checkIfFilled(Board.Cells[i][1:clearable_columns]) {
			for j := i; j > 0; j-- {
				Board.Cells[j] = Board.Cells[j-1]
			}
			Board.Cells[0] = []byte{254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254}
			i++
		}
	}
	return res
}

func (b *board) GetFilled() []int {
	var res []int
	for i := tall; i >= 0; i-- {
		if checkIfFilled(Board.Cells[i][1:clearable_columns]) {
			res = append(res, i)
		}
	}
	return res
}

func checkIfFilled(line []byte) bool {
	for _, v := range line {
		if v == 0 {
			return false
		}
	}
	return true
}

func (b *board) FillBoard() {
	for ix, row := range Board.Cells {
		for sub_ix, _ := range row {
			Board.Cells[ix][sub_ix] = byte(rand.Intn(5))
		}
	}
}
