package audio

import (
	lib "github.com/mbuechmann/terminalblaster/internal/library"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

var currentTrack *mix.Music
var currentChunk *mix.Chunk
var currentIndex int
var trackList []*lib.Track
var playing bool

func init() {
	sdl.Init(sdl.INIT_AUDIO)
}

// SetTracks sets the tracks to be played and the index of the first track to be
// played.
func SetTracks(tracks []*lib.Track, index int) (err error) {
	trackList = tracks
	currentIndex = index

	if err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 8192); err != nil {
		return
	}
	currentTrack, err = mix.LoadMUS(tracks[index].Path)

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

	currentTrack.Free()
	mix.CloseAudio()
	playing = false

	if currentIndex < len(trackList)-1 {
		SetTracks(trackList, currentIndex+1)
		go Play()
	}
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
