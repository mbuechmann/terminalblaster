package library

// Track represents one file of music and its metadata.
type Track struct {
	Path        string
	Title       string
	AlbumTitle  string
	Album       *Album
	Artist      string
	AlbumArtist string
	TrackNumber int
}
