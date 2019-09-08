package SDL

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type Background struct {
	Image *sdl.Texture
	Music *mix.Music
}

var Backgrounds map[int]Background

func LoadBackgrounds(w *sdl.Window, r *sdl.Renderer) {
	Backgrounds = make(map[int]Background)
	Backgrounds[0] = Background{
		Image: GetTexture(w, r, "backgrounds/sky.png"),
		Music: LoadMusic("sound/music/happy.mp3"),
	}
	Backgrounds[1] = Background{
		Image: GetTexture(w, r, "backgrounds/City1.png"),
		Music: LoadMusic("sound/music/happy.mp3"),
	}
	Backgrounds[2] = Background{
		Image: GetTexture(w, r, "backgrounds/Battleground1.png"),
		Music: LoadMusic("sound/music/happy.mp3"),
	}
	Backgrounds[3] = Background{
		Image: GetTexture(w, r, "backgrounds/City2.png"),
		Music: LoadMusic("sound/music/happy.mp3"),
	}
	Backgrounds[4] = Background{
		Image: GetTexture(w, r, "backgrounds/Battleground2.png"),
		Music: LoadMusic("sound/music/happy.mp3"),
	}
	Backgrounds[5] = Background{
		Image: GetTexture(w, r, "backgrounds/City3.png"),
		Music: LoadMusic("sound/music/happy.mp3"),
	}
	Backgrounds[6] = Background{
		Image: GetTexture(w, r, "backgrounds/Battleground3.png"),
		Music: LoadMusic("sound/music/happy.mp3"),
	}
	Backgrounds[7] = Background{
		Image: GetTexture(w, r, "backgrounds/City4.png"),
		Music: LoadMusic("sound/music/happy.mp3"),
	}
	Backgrounds[8] = Background{
		Image: GetTexture(w, r, "backgrounds/Battleground4.png"),
		Music: LoadMusic("sound/music/happy.mp3"),
	}
}
