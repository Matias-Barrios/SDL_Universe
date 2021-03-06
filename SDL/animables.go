package SDL

import "github.com/veandco/go-sdl2/sdl"

var BeamTextures []*sdl.Texture
var YouLoseTextures []*sdl.Texture

func LoadAnimations(w *sdl.Window, r *sdl.Renderer) {
	YouLoseTextures = []*sdl.Texture{
		GetTexture(w, r, "animations/youlose/youlose_1.png", 255),
	}
	BeamTextures = []*sdl.Texture{
		GetTexture(w, r, "animations/beam/beam1.png", 255),
		GetTexture(w, r, "animations/beam/beam2.png", 255),
		GetTexture(w, r, "animations/beam/beam3.png", 255),
		GetTexture(w, r, "animations/beam/beam4.png", 255),
	}
}
