package audio

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func play(file string) {
	c := make(chan bool)

	sdl.Init(sdl.INIT_AUDIO)

	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096*8); err != nil {
		panic(err)
	}

	mus, err := mix.LoadMUS(file)
	if err != nil {
		panic(err)
	}
	if err := mus.Play(1); err != nil {
		panic(err)
	}

	mix.HookMusicFinished(func() {
		c <- true
	})

	<-c

	mix.CloseAudio()
}
