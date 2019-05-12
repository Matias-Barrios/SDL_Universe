package SDL

import (
	"log"
	"os"

	"github.com/Matias-Barrios/SDL_Universe/definitions"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func InitSDL() (*sdl.Window, *sdl.Renderer, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return &sdl.Window{}, &sdl.Renderer{}, err
	}
	window, err := sdl.CreateWindow("Universe", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(definitions.Screen.Width), int32(definitions.Screen.Height), sdl.WINDOW_SHOWN)
	if err != nil {
		return &sdl.Window{}, &sdl.Renderer{}, err
	}
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return &sdl.Window{}, &sdl.Renderer{}, err
	}
	return window, renderer, nil
}

// Get a texture from a window, a renderer and a path
func GetTexture(w *sdl.Window, r *sdl.Renderer, path string) *sdl.Texture {
	surface, err := w.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	surfaceImg, err := img.Load(path)
	if err != nil {
		log.Fatalf("Failed to load PNG: %s\n", err)
		os.Exit(4)
	}
	textureImg, err := r.CreateTextureFromSurface(surfaceImg)
	if err != nil {
		log.Fatalf("Failed to create texture: %s\n", err)
		os.Exit(5)
	}
	surfaceImg.Free()
	return textureImg
}

// DrawStuff : Draw shit
func DrawStuff(r *sdl.Renderer, t *sdl.Texture, posx int32, posy int32, width int32, height int32) {
	r.Copy(t, nil, &sdl.Rect{posx, posy, width, width})
}
