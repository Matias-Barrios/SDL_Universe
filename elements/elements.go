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
		int32(definitions.PointsToRatioH(330)),
		int32(definitions.PointsToRatioV(100)),
		int32(definitions.PointsToRatioH(150)),
		int32(definitions.PointsToRatioV(45)),
	)
	SDL.DrawStuff(r,
		SDL.Messages_Textures["points_bar"],
		int(definitions.PointsToRatioH(330)),
		int(definitions.PointsToRatioV(150)),
		int(definitions.PointsToRatioH(150)),
		int(definitions.PointsToRatioV(140)))

	var pWidth = float64(definitions.Screen.BlockSizeW) * 0.50
	var pHeight = float64(definitions.Screen.BlockSizeH) * 0.50
	for ix, row := range p.Shape[p.Spin] {
		for sub_ix, val := range row {
			if val != 0 {
				SDL.DrawStuff(r,
					SDL.Block_Textures[SDL.Translate(val)],
					int(sub_ix*int(pWidth))+definitions.PointsToRatioH(340),
					int(ix*int(pHeight))+definitions.PointsToRatioV(190),
					int(pWidth),
					int(pHeight))
			}
		}
	}
}

func PointsBar(r *sdl.Renderer) {
	DrawText("Points", FONTS["8bitw"], r, sdl.Color{19, 109, 220, 255},
		int32(definitions.PointsToRatioH(330)),
		int32(definitions.PointsToRatioV(300)),
		int32(definitions.PointsToRatioH(170)),
		int32(definitions.PointsToRatioV(45)),
	)
	SDL.DrawStuff(r,
		SDL.Messages_Textures["points_bar"],
		int(definitions.PointsToRatioH(330)),
		int(definitions.PointsToRatioV(350)),
		int(definitions.PointsToRatioH(170)),
		int(definitions.PointsToRatioH(70)))

	DrawNumber(strconv.Itoa(definitions.Game.Points), FONTS["8bitw"], r, sdl.Color{255, 255, 255, 255},
		int32(definitions.PointsToRatioH(475)),
		int32(definitions.PointsToRatioV(370)),
		int32(definitions.PointsToRatioH(10)),
		int32(definitions.PointsToRatioV(45)))
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
	r.Copy(t, nil, &sdl.Rect{x, y, width, height})
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
	r.Copy(t, nil, &sdl.Rect{int32(int(x) - int(width)*len(text)), y, int32(int(width) * len(text)), height})
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
