package library

import "strings"

// AlbumList is a list of Albums.
type AlbumList []*Album

// Get returns the album with the matching name.
func (a AlbumList) Get(title string) *Album {
	for _, a := range a {
		if a.Title == title {
			return a
		}
	}
	return nil
}

func (a AlbumList) Len() int      { return len(a) }
func (a AlbumList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a AlbumList) Less(i, j int) bool {
	return strings.Compare(
		strings.ToLower(a[i].Title),
		strings.ToLower(a[j].Title),
	) < 0
}
