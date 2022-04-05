package audio

import (
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/speaker"

	lib "github.com/mbuechmann/terminalblaster/internal/library"
)

const bufferSize = 100

// TrackChan communicates which track gets played.
var TrackChan = make(chan *lib.Track)

// ErrorChan communicates errors during playback.
var ErrorChan = make(chan error)

var currentIndex int
var trackList []*lib.Track

var ctrl = &beep.Ctrl{}
var streamer beep.StreamSeekCloser

// SetTracks sets the tracks to be played and the index of the first track to be
// played.
func SetTracks(tracks []*lib.Track, index int) {
	trackList = tracks
	currentIndex = index
}

// Play plays the current track.
func Play() {
	// clean up if another song is already playing
	if streamer != nil {
		_ = streamer.Close()
	}

	// open file and create streamer
	track := trackList[currentIndex]
	f, err := os.Open(track.Path)
	if err != nil {
		ErrorChan <- err
	}

	var format beep.Format
	streamer, format, err = flac.Decode(f)
	if err != nil {
		ErrorChan <- err
	}

	// set up new speaker and ctrl and play stream
	speaker.Clear()
	ctrl.Paused = false
	ctrl.Streamer = streamer

	done := make(chan bool)
	speaker.Init(format.SampleRate, bufferSize)
	speaker.Play(beep.Seq(ctrl, beep.Callback(func() {
		done <- true
	})))
	<-done

	streamer.Close()
}

// Toggle pauses when playing and plays when paused.
func Toggle() {
	speaker.Lock()
	ctrl.Paused = !ctrl.Paused
	speaker.Unlock()
}
