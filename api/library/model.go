package library

import (
	"git.maych.in/thunderbottom/lastfm-go"
)

// Library represents a structure to help query the `library` LastFM API functions.
type Library struct {
	api      *lastfm.Client
	Username string
}

type libraryArtists struct {
	Artists struct {
		Artist []struct {
			Image []struct {
				Size string `json:"size"`
				Text string `json:"#text"`
			} `json:"image"`
			Mbid       string `json:"mbid"`
			Name       string `json:"name"`
			Playcount  string `json:"playcount"`
			Streamable string `json:"streamable"`
			Tagcount   string `json:"tagcount"`
			URL        string `json:"url"`
		} `json:"artist"`
		Attr struct {
			Page       string `json:"page"`
			PerPage    string `json:"perPage"`
			Total      string `json:"total"`
			TotalPages string `json:"totalPages"`
			User       string `json:"user"`
		} `json:"@attr"`
	} `json:"artists"`
}
