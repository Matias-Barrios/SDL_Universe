package board

import (
	"fmt"

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
	X: definitions.PointsToRatioH(20),
	Y: definitions.PointsToRatioV(20),
	Cells: [][]byte{
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	},
}

// this is the tall of the board minus the firs row
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
				x := definitions.PointsToRatioH(int(float64(definitions.Screen.Width) * 0.10))
				y := definitions.PointsToRatioV(int(float64(definitions.Screen.Height) * 0.10))
				width := definitions.PointsToRatioH(int(float64(definitions.Screen.Width)-float64(definitions.Screen.Width)*0.20)) - x
				height := definitions.PointsToRatioV(int(float64(definitions.Screen.Height)-float64(definitions.Screen.Height)*0.20)) - y
				fmt.Println([]int{
					int(x),
					int(y),
					int(width),
					int(height),
				})
				SDL.DrawStuff(r,
					SDL.Messages_Textures["greyblock"],
					definitions.PointsToRatioH(x),
					definitions.PointsToRatioV(y),
					definitions.PointsToRatioV(width),
					definitions.PointsToRatioV(height))
				definitions.Game.Running = false
			}
		}
	}

}

func (b *board) ClearLines() {
	for i := tall; i >= 0; i-- {
		if checkIfFilled(Board.Cells[i][1:clearable_columns]) {
			for j := i; j > 0; j-- {
				for k := 1; k < clearable_columns; k++ {
					Board.Cells[j][k] = Board.Cells[j-1][k]
				}
			}
			Board.Cells[0] = []byte{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
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
