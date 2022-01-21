package geo

import (
	"strconv"

	"git.maych.in/thunderbottom/lastfm-go"
)

// GetTopArtists fetches the most popular artists on LastFM by country.
func (g *Geo) GetTopArtists(page int) (gta *geoTopArtists, err error) {
	params := map[string]string{
		"country": g.Country,
		"limit":   g.api.GetLimit(),
		"page":    strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "geo.gettopartists",
		Params:   params,
		Response: &gta,
		Type:     "GET",
	}
	err = g.api.Request(p)

	return
}

// GetTopTracks fetches the most popular tracks in the last week on LastFM by country.
func (g *Geo) GetTopTracks(location string, page int) (gtt *geoTopTracks, err error) {
	params := map[string]string{
		"country":  g.Country,
		"limit":    g.api.GetLimit(),
		"location": location,
		"page":     strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "geo.gettoptracks",
		Params:   params,
		Response: &gtt,
		Type:     "GET",
	}
	err = g.api.Request(p)

	return
}

// New returns an instance of the `geo` API endpoint functions for LastFM.
func New(client *lastfm.Client, country string, limit int) (geo *Geo) {
	geo = &Geo{
		api:     client,
		Country: country,
	}
	return
}
