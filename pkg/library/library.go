package library

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/dhowden/tag"
)

// Artists is a list of all read Artists sorted alphabetically by artist names.
var Artists = ArtistList{}

var artistMap map[string]*Artist
var albumMap map[string]*AlbumList

var unreadableDirs = []string{}
var unreadableFiles = []string{}

var trackChan = make(chan *Track)

// Load reads the dir for the given file path and scans recursively for audio
// files. It returns a chan which provides loaded tracks.
func Load(path string) (chan *Track, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	go func() {
		artistMap = map[string]*Artist{}
		albumMap = map[string]*AlbumList{}

		filepath.Walk(path, walk)

		Artists = make(ArtistList, len(artistMap))
		i := 0
		for _, artist := range artistMap {
			Artists[i] = artist
			i++
			sort.Sort(artist.Albums)
		}
		sort.Sort(Artists)

		close(trackChan)

		artistMap = nil
		albumMap = nil
	}()

	return trackChan, nil
}

func walk(path string, info os.FileInfo, err error) error {
	if err != nil && info.IsDir() {
		unreadableDirs = append(unreadableDirs, path)
		return filepath.SkipDir
	}
	if err != nil && !info.IsDir() {
		unreadableFiles = append(unreadableFiles, path)
		return nil
	}

	f, err := os.Open(path)
	if err != nil {
		unreadableFiles = append(unreadableFiles, path)
		return nil
	}
	m, err := tag.ReadFrom(f)
	if err != nil {
		unreadableFiles = append(unreadableFiles, path)
		return nil
	}

	albumTitle := m.Album()
	albumArtist := m.AlbumArtist()
	if albumArtist == "" {
		albumArtist = m.Artist()
	}

	trackNum, _ := m.Track()
	track := &Track{
		Path:        path,
		Title:       m.Title(),
		Album:       albumTitle,
		Artist:      m.Artist(),
		AlbumArtist: albumArtist,
		TrackNumber: trackNum,
	}

	album := getAlbum(albumArtist, albumTitle)
	album.Tracks = append(album.Tracks, track)

	trackChan <- track

	return nil
}

func getArtist(name string) *Artist {
	artist, ok := artistMap[name]
	if !ok {
		artist = &Artist{Name: name, Albums: AlbumList{}}
		artistMap[name] = artist
	}

	return artist
}

func getAlbum(artistName, albumTitle string) *Album {
	artist := getArtist(artistName)
	album := artist.Albums.Get(albumTitle)
	if album == nil {
		album = &Album{Title: albumTitle, Tracks: []*Track{}}
		artist.Albums = append(artist.Albums, album)
	}

	return album
}
