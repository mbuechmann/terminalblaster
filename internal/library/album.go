package library

// Album represents a collection of tracks.
type Album struct {
	Title  string
	Tracks []*Track
}

// AddTrack adds the given track to the album.
func (a *Album) AddTrack(t *Track) {
	a.Tracks = append(a.Tracks, t)
	t.Album = a
}

// TrackIndex returns the index of the given track or -1 when the given track
// cannot be found on the album list.
func (a *Album) TrackIndex(track *Track) int {
	for i, t := range a.Tracks {
		if track == t {
			return i
		}
	}
	return -1
}
