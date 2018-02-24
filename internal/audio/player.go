package audio

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

var currentTrack *mix.Music
var playing bool

func init() {
	sdl.Init(sdl.INIT_AUDIO)
}

// Load loads the given file.
func Load(file string) (err error) {
	if err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		return
	}

	currentTrack, err = mix.LoadMUS(file)

	return
}

// Play plays the audio file at the given path.
func Play() {
	if err := currentTrack.Play(1); err != nil {
		panic(err)
	}

	playing = true

	c := make(chan bool)
	mix.HookMusicFinished(func() {
		c <- true
	})

	<-c

	mix.CloseAudio()
	playing = false
}

// Toggle pause when playing and plays when pausing.
func Toggle() {
	if currentTrack != nil {
		if playing {
			mix.PauseMusic()
		} else {
			mix.ResumeMusic()
		}
		playing = !playing
	}
}
