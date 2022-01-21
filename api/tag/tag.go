package tag

import (
	"strconv"

	"git.maych.in/thunderbottom/lastfm-go"
)

// GetInfo fetches metadata for the provided tag from LastFM.
func (t *Tag) GetInfo(tag, lang string) (ti *tagInfo, err error) {
	params := map[string]string{
		"lang": lang,
		"tag":  tag,
	}
	p := &lastfm.Provider{
		Method:   "tag.getinfo",
		Params:   params,
		Response: &ti,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// GetSimilar searches for similar tags from LastFM,
// ranked by similarity and based on listening data.
func (t *Tag) GetSimilar(tag string) (ts *tagSimilar, err error) {
	params := map[string]string{
		"tag": tag,
	}
	p := &lastfm.Provider{
		Method:   "tag.getsimilar",
		Params:   params,
		Response: &ts,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// GetTopAlbums fetches top albums from LastFM
// tagged by the provided tag, ordered by tag count.
func (t *Tag) GetTopAlbums(tag string, page int) (tta *tagTopAlbums, err error) {
	params := map[string]string{
		"limit": t.api.GetLimit(),
		"page":  strconv.Itoa(page),
		"tag":   tag,
	}
	p := &lastfm.Provider{
		Method:   "tag.gettopalbums",
		Params:   params,
		Response: &tta,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// GetTopArtists fetches top artists from LastFM
// tagged by the provided tag, ordered by tag count.
func (t *Tag) GetTopArtists(tag string, page int) (tta *tagTopArtists, err error) {
	params := map[string]string{
		"limit": t.api.GetLimit(),
		"page":  strconv.Itoa(page),
		"tag":   tag,
	}
	p := &lastfm.Provider{
		Method:   "tag.gettopartists",
		Params:   params,
		Response: &tta,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// GetTopTags fetches top tags from LastFM based on popularity.
func (t *Tag) GetTopTags() (tt *tagTopTags, err error) {
	p := &lastfm.Provider{
		Method:   "tag.gettoptags",
		Params:   map[string]string{},
		Response: &tt,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// GetTopTracks fetches top tracks from LastFM
// tagged by the provided tag, ordered by tag count.
func (t *Tag) GetTopTracks(tag string, page int) (ttt *tagTopTracks, err error) {
	params := map[string]string{
		"limit": t.api.GetLimit(),
		"page":  strconv.Itoa(page),
		"tag":   tag,
	}
	p := &lastfm.Provider{
		Method:   "tag.gettoptracks",
		Params:   params,
		Response: &ttt,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// GetWeeklyChartList fetches a list of available charts from LastFM
// for the provided tag, expressed as date ranges in unixtime.
func (t *Tag) GetWeeklyChartList(tag string) (twc *tagWeeklyChartList, err error) {
	params := map[string]string{
		"tag": tag,
	}
	p := &lastfm.Provider{
		Method:   "tag.getweeklychartlist",
		Params:   params,
		Response: &twc,
		Type:     "GET",
	}
	err = t.api.Request(p)

	return
}

// New returns an instance of the `tag` API endpoint functions for LastFM.
func New(client *lastfm.Client) (tag *Tag) {
	tag = &Tag{
		api: client,
	}
	return
}
