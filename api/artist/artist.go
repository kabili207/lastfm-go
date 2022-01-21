package artist

import (
	"fmt"
	"strconv"
	"strings"

	"git.maych.in/thunderbottom/lastfm-go"
)

// AddTags adds tags to an artist on LastFM using
// a list of user supplied tags
func (a *Artist) AddTags(artist string, tags []string) (err error) {
	if len(tags) > 10 {
		return fmt.Errorf(`AddTags limit exceeded for artist %s. Maximum Tags Allowed: 10`, artist)
	}

	params := map[string]string{
		"artist": artist,
		"tags":   strings.Join(tags, ","),
	}

	p := &lastfm.Provider{
		Method:   "artist.addtags",
		Params:   params,
		Response: nil,
		Type:     "POST",
	}
	err = a.api.Request(p)

	return
}

// GetCorrection fetches canonical artist details from LastFM
// for the provided artist
func (a *Artist) GetCorrection(artist string) (ac *artistCorrection, err error) {
	params := map[string]string{
		"artist": artist,
	}
	p := &lastfm.Provider{
		Method:   "artist.getcorrection",
		Params:   params,
		Response: &ac,
		Type:     "GET",
	}
	err = a.api.Request(p)

	return
}

// GetInfo fetches artist metadata from LastFM using artist name
// or MBID (MusicBrainz ID)
//
// language needs to be an ISO 639, alpha-2 encoded string (default: en)
func (a *Artist) GetInfo(artist, mbid, lang string) (ai *artistInfo, err error) {
	if lang == "" {
		lang = "en"
	}

	params := map[string]string{
		"artist":      artist,
		"autocorrect": a.autocorrect,
		"lang":        lang,
		"mbid":        mbid,
		"username":    a.Username,
	}
	p := &lastfm.Provider{
		Method:   "artist.getinfo",
		Params:   params,
		Response: &ai,
		Type:     "GET",
	}
	err = a.api.Request(p)

	return
}

// GetSimilar fetches similar artists from LastFM for the provided artist
// or MBID (MusicBrainz ID)
func (a *Artist) GetSimilar(artist, mbid string) (as *artistSimilar, err error) {
	params := map[string]string{
		"artist":      artist,
		"autocorrect": a.autocorrect,
		"limit":       a.api.GetLimit(),
		"mbid":        mbid,
	}
	p := &lastfm.Provider{
		Method:   "artist.getsimilar",
		Params:   params,
		Response: &as,
		Type:     "GET",
	}
	err = a.api.Request(p)

	return
}

// GetTags fetches user-applied tags on an artist from LastFM for the provided
// artist name or MBID (MusicBrainz ID)
func (a *Artist) GetTags(artist, mbid string) (at *artistTags, err error) {
	params := map[string]string{
		"artist":      artist,
		"autocorrect": a.autocorrect,
		"mbid":        mbid,
		"user":        a.Username,
	}
	p := &lastfm.Provider{
		Method:   "artist.gettags",
		Params:   params,
		Response: &at,
		Type:     "GET",
	}
	err = a.api.Request(p)

	return
}

// GetTopAlbums fetches top albums for the provided artist from LastFM,
// based on artist name or MBID (MusicBrainz ID)
func (a *Artist) GetTopAlbums(artist, mbid string, page int) (ata *artistTopAlbums, err error) {
	params := map[string]string{
		"artist":      artist,
		"autocorrect": a.autocorrect,
		"mbid":        mbid,
		"limit":       a.api.GetLimit(),
		"page":        strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "artist.gettopalbums",
		Params:   params,
		Response: &ata,
		Type:     "GET",
	}
	err = a.api.Request(p)

	return
}

// GetTopTags fetches top tags for the provided artist from LastFM,
// ordered by tag count, based on artist name or MBID (MusicBrainz ID)
func (a *Artist) GetTopTags(artist, mbid string) (att *artistTopTags, err error) {
	params := map[string]string{
		"artist":      artist,
		"autocorrect": a.autocorrect,
		"mbid":        mbid,
	}
	p := &lastfm.Provider{
		Method:   "artist.gettoptags",
		Params:   params,
		Response: &att,
		Type:     "GET",
	}
	err = a.api.Request(p)

	return
}

// GetTopTracks fetches top tracks for the provided artist from LastFM,
// based on artist name or MBID (MusicBrainz ID)
func (a *Artist) GetTopTracks(artist, mbid string, page int) (att *artistTopTracks, err error) {
	params := map[string]string{
		"artist":      artist,
		"autocorrect": a.autocorrect,
		"limit":       a.api.GetLimit(),
		"mbid":        mbid,
		"page":        strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "artist.gettoptracks",
		Params:   params,
		Response: &att,
		Type:     "GET",
	}
	err = a.api.Request(p)

	return
}

// RemoveTag removes the provided user-applied tag from an artist on LastFM
func (a *Artist) RemoveTag(artist, tag string) (err error) {
	params := map[string]string{
		"artist": artist,
		"tag":    tag,
	}
	p := &lastfm.Provider{
		Method:   "artist.removetag",
		Params:   params,
		Response: nil,
		Type:     "POST",
	}
	err = a.api.Request(p)

	return
}

// Search searches for an artist on LastFM.
func (a *Artist) Search(artist string, page int) (as *artistSearch, err error) {
	params := map[string]string{
		"artist": artist,
		"limit":  a.api.GetLimit(),
		"page":   strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "artist.search",
		Params:   params,
		Response: &as,
		Type:     "GET",
	}
	err = a.api.Request(p)

	return
}

// New returns an instance of the `album` API endpoint functions for LastFM.
func New(client *lastfm.Client, username string, autocorrect bool) (artist *Artist) {
	artist = &Artist{
		api:         client,
		autocorrect: client.Bool2strint(autocorrect),
		Username:    username,
	}
	return
}
