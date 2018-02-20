package library

import "strings"

// Artist represents an artist with name and albums.
type Artist struct {
	Name   string
	Albums AlbumList
}

// ArtistList is a list of Artists.
type ArtistList []*Artist

func (a ArtistList) Len() int      { return len(a) }
func (a ArtistList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ArtistList) Less(i, j int) bool {
	return strings.Compare(
		strings.ToLower(a[i].Name),
		strings.ToLower(a[j].Name),
	) < 0
}
