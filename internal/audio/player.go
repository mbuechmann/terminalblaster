package audio

import (
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/speaker"

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

var ctrl = &beep.Ctrl{}

var speakerInitialized bool

// SetTracks sets the tracks to be played and the index of the first track to be
// played.
func SetTracks(tracks []*lib.Track, index int) {
	trackList = tracks
	currentIndex = index
}

// Play plays the current track.
func Play() {
	track := trackList[currentIndex]
	f, err := os.Open(track.Path)
	if err != nil {
		ErrorChan <- err
	}

	streamer, format, err := flac.Decode(f)
	if err != nil {
		ErrorChan <- err
	}
	defer streamer.Close()

	sr := format.SampleRate

	if !speakerInitialized {
		speaker.Init(format.SampleRate, sr.N(time.Second/10))
		speakerInitialized = true
	}

	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	speaker.Clear()
	speaker.Lock()
	ctrl.Paused = false
	ctrl.Streamer = resampled
	speaker.Unlock()

	done := make(chan bool)
	speaker.Play(beep.Seq(ctrl, beep.Callback(func() {
		done <- true
	})))
	<-done

	currentIndex = (currentIndex + 1) % len(trackList)
	Play()
}

// Toggle pause when playing and plays when pausing.
func Toggle() {
	if ctrl != nil {
		speaker.Lock()
		ctrl.Paused = !ctrl.Paused
		speaker.Unlock()
	}
}
