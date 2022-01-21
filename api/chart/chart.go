package chart

import (
	"strconv"

	"git.maych.in/thunderbottom/lastfm-go"
)

// GetTopArtists fetches the top artists chart from LastFM.
func (c *Chart) GetTopArtists(page int) (cta *chartTopArtists, err error) {
	params := map[string]string{
		"limit": c.api.GetLimit(),
		"page":  strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "chart.gettopartists",
		Params:   params,
		Response: &cta,
		Type:     "GET",
	}
	err = c.api.Request(p)

	return
}

// GetTopTags fetches the top tags chart from LastFM.
func (c *Chart) GetTopTags(page int) (ctt *chartTopTags, err error) {
	params := map[string]string{
		"limit": c.api.GetLimit(),
		"page":  strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "chart.gettoptags",
		Params:   params,
		Response: &ctt,
		Type:     "GET",
	}
	err = c.api.Request(p)

	return
}

// GetTopTracks fetches the top tracks chart from LastFM.
func (c *Chart) GetTopTracks(page int) (ctt *chartTopTracks, err error) {
	params := map[string]string{
		"limit": c.api.GetLimit(),
		"page":  strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "chart.gettoptracks",
		Params:   params,
		Response: &ctt,
		Type:     "GET",
	}
	err = c.api.Request(p)

	return
}

// New returns an instance of the `chart` API endpoint functions for LastFM.
func New(client *lastfm.Client, country string, limit int) (chart *Chart) {
	chart = &Chart{
		api: client,
	}
	return
}
