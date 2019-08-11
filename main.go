package main

import (
	"log"
	"os"

	"github.com/Matias-Barrios/SDL_Universe/board"
	"github.com/Matias-Barrios/SDL_Universe/definitions"
	"github.com/Matias-Barrios/SDL_Universe/elements"
	"github.com/Matias-Barrios/SDL_Universe/pieces"

	"github.com/Matias-Barrios/SDL_Universe/SDL"

	"github.com/veandco/go-sdl2/sdl"
)

var ANIMATIONS []*SDL.Animable
var StopMovement bool = true

func main() {
	window, renderer, err := SDL.InitSDL()
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	defer sdl.Quit()
	defer window.Destroy()
	defer renderer.Destroy()
	t := SDL.GetTexture(window, renderer, "backgrounds/sky.png")
	SDL.LoadTextures(window, renderer)
	SDL.BricksLoadTextures(window, renderer)
	// MAIN LOOP ....
	// **************************************
	//var thePiece = pieces.Pieces[pieces.RandomPiece()]
	var thePiece = pieces.Pieces["line"]
	var next = pieces.Pieces[pieces.RandomPiece()]
	ANIMATIONS = append(ANIMATIONS, &SDL.Animable{
		Posx:     0,
		Posy:     0,
		Width:    200,
		Height:   70,
		Textures: SDL.BeamTextures,
		Timings:  []int{10, 10, 10, 400},
		Tick:     0,
		Index:    0,
		Endless:  false,
		Finished: false,
		Handler: func() {
			StopMovement = false
		},
	})
	running := true
	go func() {
		for running {
			if !StopMovement {
				sdl.Delay(definitions.Game.Delay)
				thePiece.Fall(&next)
			}
		}
	}()

	for {
		// Poll for SDL events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				os.Exit(0)
				break
			case *sdl.KeyboardEvent:
				if t.Type != sdl.KEYDOWN {
					break
				}
				switch key := t.Keysym.Sym; key {
				case sdl.K_LEFT:
					thePiece.Move(-1)
				case sdl.K_RIGHT:
					thePiece.Move(1)
				case sdl.K_a:
					thePiece.SpinIt(-1)
				case sdl.K_s:
					thePiece.SpinIt(1)
				}
			}
		}
		if definitions.Game.Running {

			// Background
			SDL.DrawStuff(renderer, t, 0, 0, int(definitions.Screen.Width), int(definitions.Screen.Height))

			// Happenings
			// ***********************

			// Draw stuff
			// ***********************

			elements.NextPieceBox(renderer, next)
			board.Draw(renderer)
			thePiece.Draw(renderer)
			board.GameOver(renderer)

			// Animables
			// ************************

			for index, a := range ANIMATIONS {
				if !a.Finished {
					a.Draw(renderer)
				} else {
					ANIMATIONS = SDL.RemoveAnimationAtIndex(ANIMATIONS, index)
				}
			}
			// // Present stuff
			// // ***********************
			renderer.Present()
		} else {

		}
		sdl.Delay(1)

	}
	sdl.Delay(2000)
	// END MAIN LOOP ....
	// **************************************
}
