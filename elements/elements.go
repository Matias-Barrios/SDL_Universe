package elements

import (
	"log"
	"os"
	"strconv"

	"github.com/Matias-Barrios/SDL_Universe/SDL"
	"github.com/Matias-Barrios/SDL_Universe/definitions"
	"github.com/Matias-Barrios/SDL_Universe/pieces"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var FONTS map[string]*ttf.Font

func NextPieceBox(r *sdl.Renderer, p pieces.Piece) {
	DrawText("Next piece", FONTS["8bitw"], r, sdl.Color{19, 109, 220, 255},
		550,
		10,
		200,
		80,
	)
	var pWidth = float64(definitions.Screen.BlockSizeW) * 0.70
	var pHeight = float64(definitions.Screen.BlockSizeH) * 0.70
	for ix, row := range p.Shape[p.Spin] {
		for sub_ix, val := range row {
			if val != 0 {
				SDL.DrawStuff(r,
					SDL.Block_Textures[SDL.Translate(val)],
					int(sub_ix*int(pWidth))+540,
					int(ix*int(pHeight))+90,
					int(pWidth),
					int(pHeight))
			}
		}
	}
}

func PointsBar(r *sdl.Renderer) {
	var x int32 = 550
	var y int32 = 310
	var width int32 = 210
	var height int32 = 80
	DrawText("Points", FONTS["8bitw"], r, sdl.Color{19, 109, 220, 255},
		(x),
		(y),
		(width),
		(height),
	)
	SDL.DrawStuff(r,
		SDL.Messages_Textures["points_bar"],
		int(x),
		int(y+height+20),
		int(width*2),
		int(float64(height)*1.10))

	DrawNumber(strconv.Itoa(definitions.Game.Points), FONTS["8bitw"], r, sdl.Color{255, 255, 255, 255},
		(x + (width * 2) - 5),
		(y + height + 29),
		(15),
		int32(float64(height)*0.80))
}

func LinesBar(r *sdl.Renderer) {
	var x int32 = 550
	var y int32 = 510
	var width int32 = 210
	var height int32 = 80

	DrawText("Lines", FONTS["8bitw"], r, sdl.Color{19, 109, 220, 255},
		(x),
		(y),
		(width),
		(height),
	)
	SDL.DrawStuff(r,
		SDL.Messages_Textures["points_bar"],
		int(x),
		int(y+height+20),
		int(width*2),
		int(float64(height)*1.10))

	DrawNumber(strconv.Itoa(definitions.Game.Lines), FONTS["8bitw"], r, sdl.Color{255, 255, 255, 255},
		(x + (width * 2) - 5),
		(y + height + 29),
		(15),
		int32(float64(height)*0.80))
}

func LevelBar(r *sdl.Renderer) {
	var x int32 = 550
	var y int32 = 810
	var width int32 = 210
	var height int32 = 80

	DrawText("Level", FONTS["8bitw"], r, sdl.Color{252, 107, 3, 255},
		(x),
		(y),
		(width),
		(height),
	)
	DrawNumber(strconv.Itoa(definitions.Game.Level), FONTS["8bitw"], r, sdl.Color{150, 200, 50, 255},
		(x + (width * 2) - 50),
		(y),
		(30),
		height)
}

func LoseAnimation(r *sdl.Renderer) {
	DrawText("You lose", FONTS["8bitw"], r, sdl.Color{219, 30, 30, 255},
		(70),
		(300),
		(370),
		(100),
	)
	r.Present()
	readKey()
	sdl.Delay(1000)
}

func readKey() {
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				os.Exit(0)
				break
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN {
					switch key := t.Keysym.Sym; key {
					default:
						return
					}
				}
			}
		}
	}
}

func DrawText(text string, font *ttf.Font, r *sdl.Renderer, color sdl.Color, x int32, y int32, width int32, height int32) {
	solid, err := font.RenderUTF8Solid(text, color)
	if err != nil {
		log.Fatalln("Render Solid - ", err.Error())
	}
	defer solid.Free()
	t, err := r.CreateTextureFromSurface(solid)
	if err != nil {
		log.Fatalf("Failed to create texture: %s\n", err)
		os.Exit(5)
	}
	r.Copy(t, nil, &sdl.Rect{int32(definitions.PointsToRatioH(float64(x))),
		int32(definitions.PointsToRatioV(float64(y))),
		int32(definitions.PointsToRatioH(float64(width))),
		int32(definitions.PointsToRatioV(float64(height)))})
}

func DrawNumber(text string, font *ttf.Font, r *sdl.Renderer, color sdl.Color, x int32, y int32, width int32, height int32) {
	solid, err := font.RenderUTF8Solid(text, color)
	if err != nil {
		log.Fatalln("Render Solid - ", err.Error())
	}
	defer solid.Free()
	t, err := r.CreateTextureFromSurface(solid)
	if err != nil {
		log.Fatalf("Failed to create texture: %s\n", err)
		os.Exit(5)
	}
	r.Copy(t, nil, &sdl.Rect{int32(definitions.PointsToRatioH(float64(int(x) - (len(text) * int(width))))),
		int32(definitions.PointsToRatioV(float64(y))),
		int32(definitions.PointsToRatioH(float64((len(text) * int(width))))),
		int32(definitions.PointsToRatioV(float64(height)))})
}

func LoadFont(path string) *ttf.Font {
	if font, err := ttf.OpenFont(path, 32); err == nil {
		return font
	} else {
		log.Fatalln(err.Error())
	}
	return nil // Unreachable
}

func init() {
	if err := ttf.Init(); err != nil {
		log.Fatalf("TTF init : %s\n", err.Error())
	}
	FONTS = make(map[string]*ttf.Font)
	FONTS["normal"] = LoadFont("fonts/test.ttf")
	FONTS["8bitw"] = LoadFont("fonts/8bitw.ttf")
}
