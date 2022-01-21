package album

import (
	"fmt"
	"strconv"
	"strings"

	"git.maych.in/thunderbottom/lastfm-go"
)

// AddTags adds tags to an album on LastFM using
// a list of user supplied tags
func (a *Album) AddTags(artist, album string, tags []string) (err error) {
	if len(tags) > 10 {
		return fmt.Errorf(`AddTags limit exceeded for "%s - %s". Maximum Tags Allowed: 10`, artist, album)
	}

	params := map[string]string{
		"album":  album,
		"artist": artist,
		"tags":   strings.Join(tags, ","),
	}

	p := &lastfm.Provider{
		Method:   "album.addtags",
		Params:   params,
		Response: nil,
		Type:     "POST",
	}
	err = a.api.Request(p)

	return
}

// GetInfo fetches album metadata from LastFM using artist and
// album name, or MBID (MusicBrainz ID)
//
// lang needs to be an ISO 639 alpha-2 encoded string (default: en)
func (a *Album) GetInfo(artist, album, mbid, lang string) (ai *albumInfo, err error) {
	if lang == "" {
		lang = "en"
	}
	params := map[string]string{
		"album":       album,
		"artist":      artist,
		"autocorrect": a.autocorrect,
		"lang":        lang,
		"mbid":        mbid,
		"username":    a.Username,
	}
	p := &lastfm.Provider{
		Method:   "album.getinfo",
		Params:   params,
		Response: &ai,
		Type:     "GET",
	}
	err = a.api.Request(p)

	return
}

// GetTags fetches user-applied tags on an album from LastFM for the provided
// artist and album name, or MBID (MusicBrainz ID)
func (a *Album) GetTags(artist, album, mbid string) (at *albumTags, err error) {
	params := map[string]string{
		"album":       album,
		"artist":      artist,
		"autocorrect": a.autocorrect,
		"mbid":        mbid,
		"user":        a.Username,
	}
	p := &lastfm.Provider{
		Method:   "album.gettags",
		Params:   params,
		Response: &at,
		Type:     "GET",
	}
	err = a.api.Request(p)

	return
}

// GetTopTags fetches top tags for the provided album from LastFM,
// ordered by tag count, based on artist and album name, or MBID (MusicBrainz ID)
func (a *Album) GetTopTags(artist, album, mbid string) (att *albumTopTags, err error) {
	params := map[string]string{
		"album":       album,
		"artist":      artist,
		"autocorrect": a.autocorrect,
		"mbid":        mbid,
	}
	p := &lastfm.Provider{
		Method:   "album.gettoptags",
		Params:   params,
		Response: &att,
		Type:     "GET",
	}
	err = a.api.Request(p)

	return
}

// RemoveTag removes the provided user-applied tag from an album on LastFM
func (a *Album) RemoveTag(artist, album, tag string) (err error) {
	params := map[string]string{
		"album":  album,
		"artist": artist,
		"tag":    tag,
	}
	p := &lastfm.Provider{
		Method:   "album.removetag",
		Params:   params,
		Response: nil,
		Type:     "POST",
	}
	err = a.api.Request(p)

	return
}

// Search searches for a track by artist and album name on LastFM.
func (a *Album) Search(artist, album string, page int) (as *albumSearch, err error) {
	params := map[string]string{
		"album":  album,
		"artist": artist,
		"limit":  a.api.GetLimit(),
		"page":   strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "album.search",
		Params:   params,
		Response: as,
		Type:     "GET",
	}
	err = a.api.Request(p)

	return
}

// New returns an instance of the `album` API endpoint functions for LastFM.
func New(client *lastfm.Client, username string, autocorrect bool) (album *Album) {
	album = &Album{
		api:         client,
		autocorrect: client.Bool2strint(autocorrect),
		Username:    username,
	}
	return
}
