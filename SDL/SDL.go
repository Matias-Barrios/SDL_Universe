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
	Lose         bool
	Level        int
}

var Ctx GameContext
var Background_Overlay *sdl.Texture

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
func GetTexture(w *sdl.Window, r *sdl.Renderer, path string, alpha uint8) *sdl.Texture {
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
	surfaceImg.SetAlphaMod(alpha)
	textureImg, err := r.CreateTextureFromSurface(surfaceImg)
	if err != nil {
		log.Fatalf("Failed to create texture: %s\n", err)
		os.Exit(5)
	}
	surfaceImg.Free()
	textureImg.SetBlendMode(sdl.BLENDMODE_BLEND)
	return textureImg
}

// DrawStuff : Draw shit
func DrawStuff(r *sdl.Renderer, t *sdl.Texture, posx int, posy int, width int, height int) {
	r.Copy(t,
		nil,
		&sdl.Rect{int32(int32(definitions.PointsToRatioH(float64(posx)))),
			int32(definitions.PointsToRatioV(float64(posy))),
			int32(definitions.PointsToRatioH(float64(width))),
			int32(definitions.PointsToRatioV(float64(height)))})
}

var Block_Textures map[string]*sdl.Texture
var Messages_Textures map[string]*sdl.Texture

// Bricks belong to https://opengameart.org/content/break-some-blocks - Chromaeleon
// Thanks dude!
func BricksLoadTextures(w *sdl.Window, r *sdl.Renderer) {
	Block_Textures = make(map[string]*sdl.Texture)
	Background_Overlay = GetTexture(w, r, "backgrounds/background_overlay.png", 0)
	Block_Textures["BlueBlockFractured"] = GetTexture(w, r, "assets/BlueBlockFractured.png", 255)
	Block_Textures["BlueBlockFX"] = GetTexture(w, r, "assets/BlueBlockFX.png", 255)
	Block_Textures["BlueBlockShatteringFX"] = GetTexture(w, r, "assets/BlueBlockShatteringFX.png", 255)
	Block_Textures["gameover"] = GetTexture(w, r, "assets/gameover.png", 255)
	Block_Textures["GrayBlockFractured"] = GetTexture(w, r, "assets/GrayBlockFractured.png", 255)
	Block_Textures["GrayBlockFX"] = GetTexture(w, r, "assets/GrayBlockFX.png", 255)
	Block_Textures["GrayBlockShatteringFX"] = GetTexture(w, r, "assets/GrayBlockShatteringFX.png", 255)
	Block_Textures["GreenBlockFractured"] = GetTexture(w, r, "assets/GreenBlockFractured.png", 255)
	Block_Textures["GreenBlockFX"] = GetTexture(w, r, "assets/GreenBlockFX.png", 255)
	Block_Textures["GreenBlockShatteringFX"] = GetTexture(w, r, "assets/GreenBlockShatteringFX.png", 255)
	Block_Textures["IcyBlueBlockFractured"] = GetTexture(w, r, "assets/IcyBlueBlockFractured.png", 255)
	Block_Textures["IcyBlueBlockFX"] = GetTexture(w, r, "assets/IcyBlueBlockFX.png", 255)
	Block_Textures["IcyBlueBlockShatteringFX"] = GetTexture(w, r, "assets/IcyBlueBlockShatteringFX.png", 255)
	Block_Textures["InfectBlockFX"] = GetTexture(w, r, "assets/InfectBlockFX.png", 255)
	Block_Textures["OrangeBlockFractured"] = GetTexture(w, r, "assets/OrangeBlockFractured.png", 255)
	Block_Textures["OrangeBlockFX"] = GetTexture(w, r, "assets/OrangeBlockFX.png", 255)
	Block_Textures["OrangeBlockShatteringFX"] = GetTexture(w, r, "assets/OrangeBlockShatteringFX.png", 255)
	Block_Textures["PinkBlockFractured"] = GetTexture(w, r, "assets/PinkBlockFractured.png", 255)
	Block_Textures["PinkBlockFX"] = GetTexture(w, r, "assets/PinkBlockFX.png", 255)
	Block_Textures["PinkBlockShatteringFX"] = GetTexture(w, r, "assets/PinkBlockShatteringFX.png", 255)
	Block_Textures["PlayerBlock"] = GetTexture(w, r, "assets/PlayerBlock.png", 255)
	Block_Textures["PurpleBlockFractured"] = GetTexture(w, r, "assets/PurpleBlockFractured.png", 255)
	Block_Textures["PurpleBlockFX"] = GetTexture(w, r, "assets/PurpleBlockFX.png", 255)
	Block_Textures["PurpleBlockShatteringFX"] = GetTexture(w, r, "assets/PurpleBlockShatteringFX.png", 255)
	Block_Textures["RainbowBlockFX"] = GetTexture(w, r, "assets/RainbowBlockFX.png", 255)
	Block_Textures["RedBlockCracked"] = GetTexture(w, r, "assets/RedBlockCracked.png", 255)
	Block_Textures["RedBlockFractured"] = GetTexture(w, r, "assets/RedBlockFractured.png", 255)
	Block_Textures["RedBlockFX"] = GetTexture(w, r, "assets/RedBlockFX.png", 255)
	Block_Textures["RestoreBlockFX"] = GetTexture(w, r, "assets/RestoreBlockFX.png", 255)
	Block_Textures["SpecialBlockFX"] = GetTexture(w, r, "assets/SpecialBlockFX.png", 255)
	Block_Textures["YellowBlockFractured"] = GetTexture(w, r, "assets/YellowBlockFractured.png", 255)
	Block_Textures["YellowBlockFX"] = GetTexture(w, r, "assets/YellowBlockFX.png", 255)
	Block_Textures["pipe_corner_bottom_left"] = GetTexture(w, r, "assets/pipe_corner_bottom_left.png", 255)
	Block_Textures["pipe_corner_bottom_right"] = GetTexture(w, r, "assets/pipe_corner_bottom_right.png", 255)
	Block_Textures["pipe_corner_top_left"] = GetTexture(w, r, "assets/pipe_corner_top_left.png", 255)
	Block_Textures["pipe_corner_top_right"] = GetTexture(w, r, "assets/pipe_corner_top_right.png", 255)
	Block_Textures["pipe_horizontal"] = GetTexture(w, r, "assets/pipe_horizontal.png", 255)
	Block_Textures["pipe_vertical"] = GetTexture(w, r, "assets/pipe_vertical.png", 255)
	Block_Textures["darkgray"] = GetTexture(w, r, "assets/darkgray.png", 255)

}

func LoadTextures(w *sdl.Window, r *sdl.Renderer) {
	Messages_Textures = make(map[string]*sdl.Texture)
	Messages_Textures["frame"] = GetTexture(w, r, "assets/frame.png", 255)
	Messages_Textures["points_bar"] = GetTexture(w, r, "assets/points_bar.png", 255)
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
