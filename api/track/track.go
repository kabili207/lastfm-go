package track

import (
	"fmt"
	"strconv"
	"strings"

	"git.maych.in/thunderbottom/lastfm-go"
)

// AddTags adds tags to a track on LastFM using
// a list of user supplied tags
func (t *Track) AddTags(artist, track string, tags []string) (err error) {
	params := map[string]string{
		"artist": artist,
		"tags":   strings.Join(tags, ","),
		"track":  track,
	}
	p := &lastfm.Provider{
		Method:   "track.addtags",
		Params:   params,
		Response: nil,
		Type:     "POST",
	}
	err = t.api.Request(p)

	return
}

// GetCorrection fetches canonical track details from LastFM
// for the provided artist and track
func (t *Track) GetCorrection(artist, track string) (tc *trackCorrection, err error) {
	params := map[string]string{
		"artist": artist,
		"track":  track,
	}
	p := &lastfm.Provider{
		Method:   "track.getcorrection",
		Params:   params,
		Response: &tc,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// GetInfo fetches track metadata from LastFM using artist and
// track name, or MBID (MusicBrainz ID)
func (t *Track) GetInfo(artist, track, mbid string) (ti *trackInfo, err error) {
	params := map[string]string{
		"artist":      artist,
		"autocorrect": t.autocorrect,
		"mbid":        mbid,
		"track":       track,
		"username":    t.Username,
	}
	p := &lastfm.Provider{
		Method:   "track.getinfo",
		Params:   params,
		Response: &ti,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// GetSimilar fetches similar tracks from LastFM for the provided artist
// and track name, or MBID (MusicBrainz ID)
func (t *Track) GetSimilar(artist, track, mbid string) (ts *trackSimilar, err error) {
	params := map[string]string{
		"artist":      artist,
		"autocorrect": t.autocorrect,
		"limit":       t.api.GetLimit(),
		"mbid":        mbid,
		"track":       track,
	}
	p := &lastfm.Provider{
		Method:   "track.getsimilar",
		Params:   params,
		Response: &ts,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// GetTags fetches user-applied tags on a track from LastFM for the provided
// artist and track name, or MBID (MusicBrainz ID)
func (t *Track) GetTags(artist, track, mbid string) (tt *trackTags, err error) {
	params := map[string]string{
		"artist":      artist,
		"autocorrect": t.autocorrect,
		"mbid":        mbid,
		"track":       track,
		"user":        t.Username,
	}
	p := &lastfm.Provider{
		Method:   "track.gettags",
		Params:   params,
		Response: &tt,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// GetTopTags fetches top tags for the provided track from LastFM,
// ordered by tag count, based on artist and track name, or MBID (MusicBrainz ID)
func (t *Track) GetTopTags(artist, track, mbid string) (ttt *trackTopTags, err error) {
	params := map[string]string{
		"artist":      artist,
		"autocorrect": t.autocorrect,
		"mbid":        mbid,
		"track":       track,
	}
	p := &lastfm.Provider{
		Method:   "track.gettoptags",
		Params:   params,
		Response: &ttt,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// Love marks the track as loved for the user on LastFM
func (t *Track) Love(artist, track string) (err error) {
	params := map[string]string{
		"artist": artist,
		"track":  track,
	}
	p := &lastfm.Provider{
		Method:   "track.love",
		Params:   params,
		Response: nil,
		Type:     "POST",
	}
	err = t.api.Request(p)

	return
}

// RemoveTag removes the provided user-applied tag from a track on LastFM
func (t *Track) RemoveTag(artist, track, tag string) (err error) {
	params := map[string]string{
		"artist": artist,
		"tag":    tag,
		"track":  track,
	}
	p := &lastfm.Provider{
		Method:   "track.removetag",
		Params:   params,
		Response: nil,
		Type:     "POST",
	}
	err = t.api.Request(p)

	return
}

// Scrobble adds a track-play to the user's profile on LastFM
// for each track in the provided scrobble list.
//
// The scrobble list needs to be a slice of the Scrobble struct.
func (t *Track) Scrobble(scrobbleList []lastfm.Scrobble) (ts *trackScrobble, err error) {
	params := map[string]string{}
	for idx, scrobble := range scrobbleList {
		if scrobble.Artist == "" || scrobble.Track == "" || scrobble.Timestamp <= 0 {
			return nil, fmt.Errorf("%v: Artist, Track, and Timestamp are mandatory for scrobbling", idx)
		}
		params[fmt.Sprintf("albumArtist[%v]", idx+1)] = scrobble.AlbumArtist
		params[fmt.Sprintf("album[%v]", idx+1)] = scrobble.Album
		params[fmt.Sprintf("artist[%v]", idx+1)] = scrobble.Artist
		params[fmt.Sprintf("chosenByUser[%v]", idx+1)] = t.api.Bool2strint(scrobble.ChosenByUser)
		params[fmt.Sprintf("context[%v]", idx+1)] = scrobble.Context
		params[fmt.Sprintf("duration[%v]", idx+1)] = strconv.FormatInt(scrobble.Duration, 10)
		params[fmt.Sprintf("mbid[%v]", idx+1)] = scrobble.MBID
		params[fmt.Sprintf("streamId[%v]", idx+1)] = scrobble.StreamID
		params[fmt.Sprintf("timestamp[%v]", idx+1)] = strconv.FormatInt(scrobble.Timestamp, 10)
		params[fmt.Sprintf("trackNumber[%v]", idx+1)] = strconv.Itoa(scrobble.TrackNumber)
		params[fmt.Sprintf("track[%v]", idx+1)] = scrobble.Track
	}
	p := &lastfm.Provider{
		Method:   "track.scrobble",
		Params:   params,
		Response: ts,
		Type:     "POST",
	}
	err = t.api.Request(p)

	return
}

// Search searches for a track by artist and track name on LastFM.
func (t *Track) Search(artist, track string, page int) (ts *trackSearch, err error) {
	params := map[string]string{
		"artist": artist,
		"limit":  t.api.GetLimit(),
		"page":   strconv.Itoa(page),
		"track":  track,
	}
	p := &lastfm.Provider{
		Method:   "track.search",
		Params:   params,
		Response: ts,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// Unlove unmarks the track as loved for the user on LastFM
func (t *Track) Unlove(artist, track string) (err error) {
	params := map[string]string{
		"artist": artist,
		"track":  track,
	}
	p := &lastfm.Provider{
		Method:   "track.unlove",
		Params:   params,
		Response: nil,
		Type:     "POST",
	}
	err = t.api.Request(p)

	return
}

// UpdateNowPlaying sets the provided track as currently playing for
// the user on LastFM.
//
// The parameter values provided for the Scrobble struct are case-sensitive.
func (t *Track) UpdateNowPlaying(scrobble lastfm.Scrobble) (tnp *trackUpdateNowPlaying, err error) {
	if scrobble.Track == "" || scrobble.Artist == "" {
		return nil, fmt.Errorf("Artist and Track name are mandatory to update now playing")
	}
	params := map[string]string{
		"album":       scrobble.Album,
		"albumArtist": scrobble.AlbumArtist,
		"artist":      scrobble.Artist,
		"context":     scrobble.Context,
		"duration":    strconv.FormatInt(scrobble.Duration, 10),
		"mbid":        scrobble.MBID,
		"track":       scrobble.Track,
		"trackNumber": strconv.Itoa(scrobble.TrackNumber),
	}
	p := &lastfm.Provider{
		Method:   "track.updatenowplaying",
		Params:   params,
		Response: tnp,
		Type:     "POST",
	}
	err = t.api.Request(p)

	return
}

// New returns an instance of the `track` API endpoint functions for LastFM.
func New(client *lastfm.Client, username string, autocorrect bool) (track *Track) {
	track = &Track{
		api:         client,
		autocorrect: client.Bool2strint(autocorrect),
		Username:    username,
	}
	return
}
