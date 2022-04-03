package audio

import (
	//"github.com/veandco/go-sdl2/mix"
	//"github.com/veandco/go-sdl2/sdl"

	lib "github.com/mbuechmann/terminalblaster/internal/library"
)

// TrackChan communicates which track gets played.
var TrackChan = make(chan *lib.Track)

// ErrorChan communicates errors during playback.
var ErrorChan = make(chan error)

//var currentTrack *mix.Music
//var currentChunk *mix.Chunk
var currentIndex int
var trackList []*lib.Track
var playing bool

func init() {
	//if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
	//	ErrorChan <- err
	//}
}

// SetTracks sets the tracks to be played and the index of the first track to be
// played.
func SetTracks(tracks []*lib.Track, index int) {
	//trackList = tracks
	//currentIndex = index
}

// Play plays the current track.
func Play() {
	//go func() {
	//	var err error
	//
	//	if err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 8192); err != nil {
	//		ErrorChan <- err
	//	}
	//
	//	if currentTrack, err = mix.LoadMUS(trackList[currentIndex].Path); err != nil {
	//		ErrorChan <- err
	//	}
	//
	//	if err := currentTrack.Play(1); err != nil {
	//		ErrorChan <- err
	//	} else {
	//		TrackChan <- trackList[currentIndex]
	//	}
	//
	//	playing = true
	//
	//	c := make(chan bool)
	//	mix.HookMusicFinished(func() {
	//		c <- true
	//	})
	//
	//	<-c
	//
	//	currentTrack.Free()
	//	mix.CloseAudio()
	//	playing = false
	//
	//	if currentIndex < len(trackList)-1 {
	//		currentIndex++
	//		Play()
	//	}
	//}()
}

// Toggle pause when playing and plays when pausing.
func Toggle() {
	//if currentTrack != nil {
	//	if playing {
	//		mix.PauseMusic()
	//	} else {
	//		mix.ResumeMusic()
	//	}
	//	playing = !playing
	//}
}
