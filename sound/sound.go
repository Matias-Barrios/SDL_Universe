package sound

if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
	panic(err)
	return
}