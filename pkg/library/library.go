package library

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/dhowden/tag"
)

// Artists is a list of all read Artists sorted alphabetically by artist names.
var Artists = []Artist{}

var artistMap = map[string]Artist{}
var artistNames = []string{}

var unreadableDirs = []string{}
var unreadableFiles = []string{}

var trackChan = make(chan *Track)

// Track represents one file of music and its metadata.
type Track struct {
	Path        string
	Title       string
	Album       string
	Artist      string
	AlbumArtist string
	TrackNumber int
}

// Album represents a collection of tracks.
type Album struct {
	Title      string
	Tracks     []Track
	TrackCount int
}

// Artist represents an artist with name and albums.
type Artist struct {
	Name   string
	Albums map[string]Album
}

// Load reads the dir for the given file path and scans recursively for audio
// files. It returns a chan which provides loaded tracks.
func Load(path string) (chan *Track, error) {
	if _, err := os.Stat(path); err != nil {
		return nil, err
	}

	go func() {
		filepath.Walk(path, walk)

		sort.Strings(artistNames)

		Artists = make([]Artist, len(artistNames))
		for i, name := range artistNames {
			Artists[i] = artistMap[name]
		}

		close(trackChan)
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

	trackNum, trackTotal := m.Track()
	track := Track{
		Path:        path,
		Title:       m.Title(),
		Album:       albumTitle,
		Artist:      m.Artist(),
		AlbumArtist: albumArtist,
		TrackNumber: trackNum,
	}

	album := getAlbum(albumArtist, albumTitle)
	album.TrackCount = trackTotal
	album.Tracks = append(album.Tracks, track)

	trackChan <- &track

	return nil
}

func getArtist(name string) Artist {
	artist, ok := artistMap[name]
	if !ok {
		artist = Artist{Name: name, Albums: map[string]Album{}}
		artistMap[name] = artist
		artistNames = append(artistNames, name)
	}

	return artist
}

func getAlbum(artistName, albumTitle string) Album {
	artist := getArtist(artistName)
	album, ok := artist.Albums[albumTitle]
	if !ok {
		album = Album{Title: albumTitle, Tracks: []Track{}}
	}
	return album
}
