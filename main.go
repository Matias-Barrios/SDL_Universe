package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Matias-Barrios/SDL_Universe/definitions"

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
	fmt.Println("Hola!!!")
	t := SDL.GetTexture(window, renderer, "assets/uni.jpeg")
	block := SDL.GetTexture(window, renderer, "assets/block.png")

	// MAIN LOOP ....
	// **************************************
	running := true
	for running {
		// Poll for SDL events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
		// Draw stuff
		SDL.DrawStuff(renderer, t, 0, 0, definitions.Screen.Width, definitions.Screen.Height)
		SDL.DrawStuff(renderer, block, 0, 0, 50, 50)

		renderer.Present()
		time.Sleep(time.Millisecond * 10)
	}

	// END MAIN LOOP ....
	// **************************************
}
