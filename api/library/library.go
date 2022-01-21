package library

import (
	"strconv"

	"git.maych.in/thunderbottom/lastfm-go"
)

// GetArtists fetches all artists in the user's library, with play counts and tag counts from LastFM.
func (l *Library) GetArtists(artist string, page int) (la *libraryArtists, err error) {
	params := map[string]string{
		"artist": artist,
		"limit":  l.api.GetLimit(),
		"page":   strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "library.getartists",
		Params:   params,
		Response: &la,
		Type:     "GET",
	}
	err = l.api.Request(p)

	return
}

// New returns an instance of the `library` API endpoint functions for LastFM.
func New(client *lastfm.Client, username string, limit int) (library *Library) {
	library = &Library{
		api:      client,
		Username: username,
	}
	return
}
