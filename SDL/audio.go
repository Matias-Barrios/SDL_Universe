package SDL

import (
	"io/ioutil"
	"log"

	"github.com/veandco/go-sdl2/mix"
)

var AUDIOS map[string]*mix.Chunk

func init() {
	AUDIOS = make(map[string]*mix.Chunk)
	AUDIOS["point_normal"] = LoadAudio("sound/effects/fuse.wav")
}

func LoadAudio(path string) *mix.Chunk {
	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Fatalf("%s\n", err.Error())
	}
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

func CloseAudio() {
	mix.CloseAudio()
}
