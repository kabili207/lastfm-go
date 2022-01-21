package geo

import (
	"git.maych.in/thunderbottom/lastfm-go"
)

// Geo represents a structure to help query the `geo` LastFM API functions.
type Geo struct {
	api     *lastfm.Client
	Country string
}

type attributes struct {
	Country    string `json:"country"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

type image struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}

type geoTopArtists struct {
	TopArtists struct {
		Artist []struct {
			Image      []image `json:"image"`
			Listeners  string  `json:"listeners"`
			Mbid       string  `json:"mbid"`
			Name       string  `json:"name"`
			Streamable string  `json:"streamable"`
			URL        string  `json:"url"`
		} `json:"artist"`
		Attributes attributes `json:"@attr"`
	} `json:"topartists"`
}

type geoTopTracks struct {
	Tracks struct {
		Attributes attributes `json:"@attr"`
		Track      []struct {
			Artist struct {
				Name string `json:"name"`
				Mbid string `json:"mbid"`
				URL  string `json:"url"`
			} `json:"artist"`
			Attributes struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
			Duration   string  `json:"duration"`
			Image      []image `json:"image"`
			Listeners  string  `json:"listeners"`
			Mbid       string  `json:"mbid"`
			Name       string  `json:"name"`
			URL        string  `json:"url"`
			Streamable struct {
				Text      string `json:"#text"`
				Fulltrack string `json:"fulltrack"`
			} `json:"streamable"`
		} `json:"track"`
	} `json:"tracks"`
}
