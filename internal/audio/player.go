package audio

import (
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/speaker"

	lib "github.com/mbuechmann/terminalblaster/internal/library"
)

const bufferSize = 100

var ctrl = &beep.Ctrl{}
var streamer beep.StreamSeekCloser

// Play plays the current track.
func Play(track *lib.Track) error {
	// clean up if another song is already playing
	if streamer != nil {
		err := streamer.Close()
		if err != nil {
			return err
		}
	}

	// open file and create streamer
	f, err := os.Open(track.Path)
	if err != nil {
		return err
	}

	var format beep.Format
	streamer, format, err = flac.Decode(f)
	if err != nil {
		return err
	}

	// set up new speaker and ctrl and play stream
	speaker.Clear()
	ctrl.Paused = false
	ctrl.Streamer = streamer

	done := make(chan bool)
	if err = speaker.Init(format.SampleRate, bufferSize); err != nil {
		return err
	}
	speaker.Play(beep.Seq(ctrl, beep.Callback(func() {
		done <- true
	})))
	<-done

	return streamer.Close()
}

// Toggle pauses when playing and plays when paused.
func Toggle() {
	speaker.Lock()
	ctrl.Paused = !ctrl.Paused
	speaker.Unlock()
}
