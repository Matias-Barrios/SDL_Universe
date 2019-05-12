package main

import (
	"fmt"
	"log"

	"github.com/Matias-Barrios/SDL_Universe/board"
	"github.com/Matias-Barrios/SDL_Universe/definitions"
	"github.com/Matias-Barrios/SDL_Universe/pieces"

	"github.com/Matias-Barrios/SDL_Universe/SDL"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	window, renderer, err := SDL.InitSDL()
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	defer sdl.Quit()
	defer window.Destroy()
	defer renderer.Destroy()
	fmt.Println(definitions.Screen.BlockSize)
	t := SDL.GetTexture(window, renderer, "assets/uni.jpeg")
	block := SDL.GetTexture(window, renderer, "assets/block.png")

	// MAIN LOOP ....
	// **************************************
	//var thePiece = pieces.Pieces[pieces.RandomPiece()]
	var thePiece = pieces.Pieces["line"]
	running := true
	go func() {
		for {
			sdl.Delay(50)
			thePiece.Fall()
		}
	}()
	for running {
		// Poll for SDL events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
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
		// Background
		SDL.DrawStuff(renderer, t, 0, 0, int32(definitions.Screen.Width), int32(definitions.Screen.Height))

		// Happenings
		// ***********************

		// Draw stuff
		// ***********************

		board.Draw(renderer, block)
		thePiece.Draw(renderer, block)
		// Present stuff
		// ***********************
		renderer.Present()

	}

	// END MAIN LOOP ....
	// **************************************
}
