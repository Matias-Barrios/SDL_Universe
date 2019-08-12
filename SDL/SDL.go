package SDL

import (
	"fmt"
	"log"
	"os"

	"github.com/Matias-Barrios/SDL_Universe/definitions"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type GameContext struct {
	ANIMATIONS   []*Animable
	StopMovement bool
	ClearLines   bool
}

var Ctx GameContext

func InitSDL() (*sdl.Window, *sdl.Renderer, error) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return &sdl.Window{}, &sdl.Renderer{}, err
	}

	mode, err := sdl.GetDesktopDisplayMode(0)
	if err != nil {
		return nil, nil, err
	}
	definitions.Screen.Width = int(float64(mode.W) / 2.3)
	definitions.Screen.Height = int(float64(mode.H) / 1.3)
	//definitions.Screen.Width = int(1024 / 2.3)
	//definitions.Screen.Height = int(800 / 1.3)

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
func DrawStuff(r *sdl.Renderer, t *sdl.Texture, posx int, posy int, width int, height int) {
	r.Copy(t, nil, &sdl.Rect{int32(posx), int32(posy), int32(width), int32(height)})
}

var Block_Textures map[string]*sdl.Texture
var Messages_Textures map[string]*sdl.Texture
var BeamTextures []*sdl.Texture

// Bricks belong to https://opengameart.org/content/break-some-blocks - Chromaeleon
// Thanks dude!
func BricksLoadTextures(w *sdl.Window, r *sdl.Renderer) {
	Block_Textures = make(map[string]*sdl.Texture)
	Block_Textures["BlueBlockFractured"] = GetTexture(w, r, "assets/BlueBlockFractured.png")
	Block_Textures["BlueBlockFX"] = GetTexture(w, r, "assets/BlueBlockFX.png")
	Block_Textures["BlueBlockShatteringFX"] = GetTexture(w, r, "assets/BlueBlockShatteringFX.png")
	Block_Textures["gameover"] = GetTexture(w, r, "assets/gameover.png")
	Block_Textures["GrayBlockFractured"] = GetTexture(w, r, "assets/GrayBlockFractured.png")
	Block_Textures["GrayBlockFX"] = GetTexture(w, r, "assets/GrayBlockFX.png")
	Block_Textures["GrayBlockShatteringFX"] = GetTexture(w, r, "assets/GrayBlockShatteringFX.png")
	Block_Textures["GreenBlockFractured"] = GetTexture(w, r, "assets/GreenBlockFractured.png")
	Block_Textures["GreenBlockFX"] = GetTexture(w, r, "assets/GreenBlockFX.png")
	Block_Textures["GreenBlockShatteringFX"] = GetTexture(w, r, "assets/GreenBlockShatteringFX.png")
	Block_Textures["IcyBlueBlockFractured"] = GetTexture(w, r, "assets/IcyBlueBlockFractured.png")
	Block_Textures["IcyBlueBlockFX"] = GetTexture(w, r, "assets/IcyBlueBlockFX.png")
	Block_Textures["IcyBlueBlockShatteringFX"] = GetTexture(w, r, "assets/IcyBlueBlockShatteringFX.png")
	Block_Textures["InfectBlockFX"] = GetTexture(w, r, "assets/InfectBlockFX.png")
	Block_Textures["OrangeBlockFractured"] = GetTexture(w, r, "assets/OrangeBlockFractured.png")
	Block_Textures["OrangeBlockFX"] = GetTexture(w, r, "assets/OrangeBlockFX.png")
	Block_Textures["OrangeBlockShatteringFX"] = GetTexture(w, r, "assets/OrangeBlockShatteringFX.png")
	Block_Textures["PinkBlockFractured"] = GetTexture(w, r, "assets/PinkBlockFractured.png")
	Block_Textures["PinkBlockFX"] = GetTexture(w, r, "assets/PinkBlockFX.png")
	Block_Textures["PinkBlockShatteringFX"] = GetTexture(w, r, "assets/PinkBlockShatteringFX.png")
	Block_Textures["PlayerBlock"] = GetTexture(w, r, "assets/PlayerBlock.png")
	Block_Textures["PurpleBlockFractured"] = GetTexture(w, r, "assets/PurpleBlockFractured.png")
	Block_Textures["PurpleBlockFX"] = GetTexture(w, r, "assets/PurpleBlockFX.png")
	Block_Textures["PurpleBlockShatteringFX"] = GetTexture(w, r, "assets/PurpleBlockShatteringFX.png")
	Block_Textures["RainbowBlockFX"] = GetTexture(w, r, "assets/RainbowBlockFX.png")
	Block_Textures["RedBlockCracked"] = GetTexture(w, r, "assets/RedBlockCracked.png")
	Block_Textures["RedBlockFractured"] = GetTexture(w, r, "assets/RedBlockFractured.png")
	Block_Textures["RedBlockFX"] = GetTexture(w, r, "assets/RedBlockFX.png")
	Block_Textures["RestoreBlockFX"] = GetTexture(w, r, "assets/RestoreBlockFX.png")
	Block_Textures["SpecialBlockFX"] = GetTexture(w, r, "assets/SpecialBlockFX.png")
	Block_Textures["YellowBlockFractured"] = GetTexture(w, r, "assets/YellowBlockFractured.png")
	Block_Textures["YellowBlockFX"] = GetTexture(w, r, "assets/YellowBlockFX.png")
	Block_Textures["pipe_corner_bottom_left"] = GetTexture(w, r, "assets/pipe_corner_bottom_left.png")
	Block_Textures["pipe_corner_bottom_right"] = GetTexture(w, r, "assets/pipe_corner_bottom_right.png")
	Block_Textures["pipe_corner_top_left"] = GetTexture(w, r, "assets/pipe_corner_top_left.png")
	Block_Textures["pipe_corner_top_right"] = GetTexture(w, r, "assets/pipe_corner_top_right.png")
	Block_Textures["pipe_horizontal"] = GetTexture(w, r, "assets/pipe_horizontal.png")
	Block_Textures["pipe_vertical"] = GetTexture(w, r, "assets/pipe_vertical.png")
	Block_Textures["darkgray"] = GetTexture(w, r, "assets/darkgray.png")

}
func LoadTextures(w *sdl.Window, r *sdl.Renderer) {
	Messages_Textures = make(map[string]*sdl.Texture)
	Messages_Textures["gameover"] = GetTexture(w, r, "assets/gameover.png")
	Messages_Textures["frame"] = GetTexture(w, r, "assets/frame.png")
	BeamTextures = []*sdl.Texture{
		GetTexture(w, r, "animations/beam/beam1.png"),
		GetTexture(w, r, "animations/beam/beam2.png"),
		GetTexture(w, r, "animations/beam/beam3.png"),
		GetTexture(w, r, "animations/beam/beam4.png"),
	}
}

func Translate(number byte) string {
	switch number {
	case 1:
		return "PinkBlockFX"
	case 2:
		return "BlueBlockFX"
	case 3:
		return "GreenBlockFX"
	case 4:
		return "IcyBlueBlockFX"
	case 5:
		return "OrangeBlockFX"
	case 6:
		return "PinkBlockFX"
	case 7:
		return "PurpleBlockFX"
	case 8:
		return "RedBlockFX"
	case 9:
		return "YellowBlockFX"
	case 10:
		return "RainbowBlockFX"
	case 251:
		return "darkgray"
	case 252:
		return "darkgray"
	case 253:
		return "darkgray"
	case 254:
		return "darkgray"
	default:
		log.Fatalf("%s\n", fmt.Errorf("Unknown identifier for texture"))
	}
	return ""
}

type Animable struct {
	Posx     int
	Posy     int
	Height   int
	Width    int
	Textures []*sdl.Texture
	Timings  []int
	Tick     int
	Index    int
	Endless  bool
	Finished bool
	Handler  func()
}

func (a *Animable) Draw(r *sdl.Renderer) {
	DrawStuff(r, a.Textures[a.Index], a.Posx, a.Posy, a.Width, a.Height)
	if a.Tick > a.Timings[a.Index] {
		a.Index++
		if a.Index > len(a.Timings)-1 {
			a.Index = 0
			if a.Endless == false {
				a.Finished = true
				if a.Handler != nil {
					a.Handler()
				}
			}
		}
		a.Tick = 0
	} else {
		a.Tick++
	}

}

func RemoveAnimationAtIndex(slice []*Animable, s int) []*Animable {
	return append(slice[:s], slice[s+1:]...)
}
