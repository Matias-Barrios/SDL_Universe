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
	X: (definitions.Screen.Width / 2) - (definitions.Screen.BlockSize * 5),
	Y: 0 - (20 * definitions.Screen.Height / 100),
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
					int32((sub_ix*definitions.Screen.BlockSize)+Board.X),
					int32((ix*definitions.Screen.BlockSize)+Board.Y),
					int32(definitions.Screen.BlockSize),
					int32(definitions.Screen.BlockSize))
			}
		}
	}

}

func GameOver(r *sdl.Renderer) {
	for ix, row := range Board.Cells {
		for sub_ix, val := range row {
			if ix > 7 {
				return
			} else if sub_ix == 0 || sub_ix == 19 {
				continue
			} else if val != 0 {
				fmt.Printf("ix : %d sub_ix : %d Value : %d\n", ix, sub_ix, Board.Cells[ix][sub_ix])
				SDL.DrawStuff(r,
					SDL.Messages_Textures["gameover"],
					int32(float64(definitions.Screen.Width)*.10),
					int32(definitions.Screen.Height-int(float64(definitions.Screen.Height)*.10)),
					int32(definitions.Screen.Width-int(float64(definitions.Screen.Width)*.10)),
					int32(definitions.Screen.Height-int(float64(definitions.Screen.Height)*.10)))
			}
			definitions.Game.Running = false
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
