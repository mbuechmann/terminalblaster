package audio

import (
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/speaker"

	lib "github.com/mbuechmann/terminalblaster/internal/library"
)

const quality = 10

// TrackChan communicates which track gets played.
var TrackChan = make(chan *lib.Track)

// ErrorChan communicates errors during playback.
var ErrorChan = make(chan error)

var currentIndex int
var trackList []*lib.Track

var ctrl = &beep.Ctrl{}

func init() {
	speaker.Init(44100, 100)
}

// SetTracks sets the tracks to be played and the index of the first track to be
// played.
func SetTracks(tracks []*lib.Track, index int) {
	trackList = tracks
	currentIndex = index
}

// Play plays the current track.
func Play() {
	// open file and create streamer
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

	// set up ctrl and play stream
	speaker.Clear()
	speaker.Lock()
	ctrl.Paused = false
	ctrl.Streamer = beep.Resample(quality, format.SampleRate, format.SampleRate, streamer)
	speaker.Unlock()

	done := make(chan bool)
	speaker.Play(beep.Seq(ctrl, beep.Callback(func() {
		done <- true
	})))
	<-done

	// advance to next track or loop
	currentIndex = (currentIndex + 1) % len(trackList)
	Play()
}

// Toggle pauses when playing and plays when paused.
func Toggle() {
	speaker.Lock()
	ctrl.Paused = !ctrl.Paused
	speaker.Unlock()
}
