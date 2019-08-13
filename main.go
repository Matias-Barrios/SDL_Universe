package main

import (
	"log"
	"os"
	"sync"

	"github.com/Matias-Barrios/SDL_Universe/board"
	"github.com/Matias-Barrios/SDL_Universe/definitions"
	"github.com/Matias-Barrios/SDL_Universe/elements"
	"github.com/Matias-Barrios/SDL_Universe/pieces"

	"github.com/Matias-Barrios/SDL_Universe/SDL"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Setting Game context
	SDL.Ctx.ANIMATIONS = make([]*SDL.Animable, 0, 100)
	SDL.Ctx.StopMovement = false
	SDL.Ctx.ClearLines = true

	// Main WaitGroup
	var wg sync.WaitGroup

	window, renderer, err := SDL.InitSDL()
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		panic(err)
		return
	}

	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
		return
	}

	defer sdl.Quit()
	defer window.Destroy()
	defer renderer.Destroy()
	defer mix.CloseAudio()
	t := SDL.GetTexture(window, renderer, "backgrounds/sky.png")

	SDL.LoadTextures(window, renderer)
	SDL.BricksLoadTextures(window, renderer)
	// MAIN LOOP ....
	// **************************************
	//var thePiece = pieces.Pieces[pieces.RandomPiece()]
	var thePiece = pieces.Pieces["line"]
	var next = pieces.Pieces[pieces.RandomPiece()]

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
				case sdl.K_DOWN:
					thePiece.Fall(&next, &SDL.Ctx)
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
			if !SDL.Ctx.StopMovement {
				thePiece.Fall(&next, &SDL.Ctx)
			}
			elements.NextPieceBox(renderer, next)
			board.Draw(renderer)
			thePiece.Draw(renderer)
			board.GameOver(renderer)

			// Animables
			// ************************
			wg.Add(len(SDL.Ctx.ANIMATIONS))
			for _, a := range SDL.Ctx.ANIMATIONS {
				go func(anim *SDL.Animable, wait *sync.WaitGroup) {
					if !anim.Finished {
						anim.Draw(renderer)
					}
					wait.Done()
				}(a, &wg)
			}
			wg.Wait()
		LOOP:
			for {
				for index, a := range SDL.Ctx.ANIMATIONS {
					if a.Finished {
						SDL.Ctx.ANIMATIONS = SDL.RemoveAnimationAtIndex(SDL.Ctx.ANIMATIONS, index)
						continue LOOP
					}
				}
				break
			}
			if SDL.Ctx.ClearLines {
				board.Board.ClearLines()
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
