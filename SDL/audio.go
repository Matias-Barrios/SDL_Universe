package SDL

import (
	"io/ioutil"
	"log"

	"github.com/veandco/go-sdl2/mix"
)

var AUDIOS map[string]*mix.Chunk
var MUSIC map[string]*mix.Music

func init() {
	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	AUDIOS = make(map[string]*mix.Chunk)
	MUSIC = make(map[string]*mix.Music)
	MUSIC["level1"] = LoadMusic("sound/music/happy.mp3")
	AUDIOS["piecedrop"] = LoadAudio("sound/effects/piece_drop.wav")
	AUDIOS["clearedLineCommon"] = LoadAudio("sound/effects/clearedLineCommon.wav")
}

func LoadAudio(path string) *mix.Chunk {
	// Load entire WAV data from file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	// Load WAV from data (memory)
	chunk, err := mix.QuickLoadWAV(data)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	return chunk
}

func LoadMusic(path string) *mix.Music {
	// Load WAV from data (memory)
	music, err := mix.LoadMUS(path)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
	return music
}

func CloseAudio() {
	mix.CloseAudio()
}

func IsPlaying(channel int) bool {
	return mix.Playing(channel) == 1
}
