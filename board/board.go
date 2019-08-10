package board

import (
	"fmt"
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

func GameOver(r *sdl.Renderer) {
	for ix, row := range Board.Cells {
		for sub_ix, val := range row {
			if ix > 7 {
				return
			} else if sub_ix < 1 || sub_ix > 10 {
				continue
			} else if val != 0 {
				x := definitions.PointsToRatioH(float64(definitions.Screen.Width) * 0.10)
				y := definitions.PointsToRatioV(float64(definitions.Screen.Height) * 0.30)
				width := definitions.PointsToRatioH(float64(definitions.Screen.Width)-float64(definitions.Screen.Width)*0.10) - x
				height := definitions.PointsToRatioV(float64(definitions.Screen.Height)-float64(definitions.Screen.Height)*0.10) - y
				fmt.Println([]int{
					x,
					y,
					width,
					height,
				})
				SDL.DrawStuff(r,
					SDL.Messages_Textures["gameover"],
					definitions.PointsToRatioH(float64(x)),
					definitions.PointsToRatioV(float64(y)),
					definitions.PointsToRatioV(float64(width)),
					definitions.PointsToRatioV(float64(height)))
				definitions.Game.Running = false
			}
		}
	}

}

func (b *board) ClearLines() {
	for i := tall; i >= 0; i-- {
		if checkIfFilled(Board.Cells[i][1:clearable_columns]) {
			for j := i; j > 0; j-- {
				Board.Cells[j] = Board.Cells[j-1]
			}
			Board.Cells[0] = []byte{254, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 254}
			i++
		}
	}
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
