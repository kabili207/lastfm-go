package chart

import (
	"git.maych.in/thunderbottom/lastfm-go"
)

// Chart represents a structure to help query the `chart` LastFM API functions.
type Chart struct {
	api *lastfm.Client
}

type attributes struct {
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

type image struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}

type chartTopTracks struct {
	Tracks struct {
		Attributes attributes `json:"@attr"`
		Track      []struct {
			Duration   string  `json:"duration"`
			Image      []image `json:"image"`
			Listeners  string  `json:"listeners"`
			Mbid       string  `json:"mbid"`
			Name       string  `json:"name"`
			Playcount  string  `json:"playcount"`
			URL        string  `json:"url"`
			Streamable struct {
				Fulltrack string `json:"fulltrack"`
				Text      string `json:"#text"`
			} `json:"streamable"`
			Artist struct {
				Mbid string `json:"mbid"`
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"artist"`
		} `json:"track"`
	} `json:"tracks"`
}

type chartTopTags struct {
	Tags struct {
		Attributes attributes `json:"@attr"`
		Tag        []struct {
			Name       string `json:"name"`
			URL        string `json:"url"`
			Reach      string `json:"reach"`
			Taggings   string `json:"taggings"`
			Streamable string `json:"streamable"`
			Wiki       struct {
			} `json:"wiki"`
		} `json:"tag"`
	} `json:"tags"`
}

type chartTopArtists struct {
	Artists struct {
		Attributes attributes `json:"@attr"`
		Artist     []struct {
			Name       string  `json:"name"`
			Playcount  string  `json:"playcount"`
			Listeners  string  `json:"listeners"`
			Mbid       string  `json:"mbid"`
			URL        string  `json:"url"`
			Streamable string  `json:"streamable"`
			Image      []image `json:"image"`
		} `json:"artist"`
	} `json:"artists"`
}
