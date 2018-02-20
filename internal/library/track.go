package library

// Track represents one file of music and its metadata.
type Track struct {
	Path        string
	Title       string
	Album       string
	Artist      string
	AlbumArtist string
	TrackNumber int
}
