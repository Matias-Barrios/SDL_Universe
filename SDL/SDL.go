package SDL

import (
	"fmt"
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

	mode, err := sdl.GetDesktopDisplayMode(0)
	if err != nil {
		return nil, nil, err
	}
	//definitions.Screen.Width = int(float64(mode.W) / 2.3)
	//definitions.Screen.Height = int(float64(mode.H) / 1.3)
	definitions.Screen.Width = int(1024.0 / 2.3)
	definitions.Screen.Height = int(800.0 / 1.3)

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

var Block_Textures map[string]*sdl.Texture
var Messages_Textures map[string]*sdl.Texture

func BricksLoadTextures(w *sdl.Window, r *sdl.Renderer) {
	Block_Textures = make(map[string]*sdl.Texture)
	Block_Textures["greyblock"] = GetTexture(w, r, "assets/greyblock.png")
	Block_Textures["greenblock"] = GetTexture(w, r, "assets/greenblock.png")
	Block_Textures["orangeblock"] = GetTexture(w, r, "assets/orangeblock.png")
	Block_Textures["purpleblock"] = GetTexture(w, r, "assets/purpleblock.png")
	Block_Textures["redblock"] = GetTexture(w, r, "assets/redblock.png")
	Block_Textures["yellowblock"] = GetTexture(w, r, "assets/yellowblock.png")
}
func LoadTextures(w *sdl.Window, r *sdl.Renderer) {
	Messages_Textures = make(map[string]*sdl.Texture)
	Messages_Textures["gameover"] = GetTexture(w, r, "assets/gameover.png")
}

func Translate(number byte) string {
	switch number {
	case 1:
		return "greyblock"
	case 2:
		return "greenblock"
	case 3:
		return "orangeblock"
	case 4:
		return "purpleblock"
	case 5:
		return "redblock"
	case 6:
		return "yellowblock"
	default:
		log.Fatalf("%s\n", fmt.Errorf("Unknown identifier for texture"))
	}
	return ""
}
