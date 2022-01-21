package tag

import (
	"git.maych.in/thunderbottom/lastfm-go"
)

// Tag represents a structure to help query the `tag` LastFM API functions.
type Tag struct {
	api *lastfm.Client
}

type artist struct {
	Name string `json:"name"`
	Mbid string `json:"mbid"`
	URL  string `json:"url"`
}

type attributes struct {
	Tag        string `json:"tag"`
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	TotalPages string `json:"totalPages"`
	Total      string `json:"total"`
}

type image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

type tagInfo struct {
	Tag struct {
		Name  string `json:"name"`
		Total int    `json:"total"`
		Reach int    `json:"reach"`
		Wiki  struct {
			Summary string `json:"summary"`
			Content string `json:"content"`
		} `json:"wiki"`
	} `json:"tag"`
}

type tagSimilar struct {
	SimilarTags struct {
		Tag []struct {
			Name       string `json:"name"`
			URL        string `json:"url"`
			Streamable int    `json:"streamable"`
		} `json:"tag"`
		Attributes struct {
			Tag string `json:"tag"`
		} `json:"@attr"`
	} `json:"similartags"`
}

type tagTopAlbums struct {
	Albums struct {
		Album []struct {
			Name       string  `json:"name"`
			Mbid       string  `json:"mbid"`
			URL        string  `json:"url"`
			Artist     artist  `json:"artist"`
			Image      []image `json:"image"`
			Attributes struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
		} `json:"album"`
		Attributes attributes `json:"@attr"`
	} `json:"albums"`
}

type tagTopArtists struct {
	TopArtists struct {
		Artist []struct {
			Name       string  `json:"name"`
			Mbid       string  `json:"mbid"`
			URL        string  `json:"url"`
			Streamable string  `json:"streamable"`
			Image      []image `json:"image"`
			Attributes struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
		} `json:"artist"`
		Attributes attributes `json:"@attr"`
	} `json:"topartists"`
}

type tagTopTags struct {
	TopTags struct {
		Attributes struct {
			Offset int `json:"offset"`
			NumRes int `json:"num_res"`
			Total  int `json:"total"`
		} `json:"@attr"`
		Tag []struct {
			Name  string `json:"name"`
			Count int    `json:"count"`
			Reach int    `json:"reach"`
		} `json:"tag"`
	} `json:"toptags"`
}

type tagTopTracks struct {
	Tracks struct {
		Track []struct {
			Name       string `json:"name"`
			Duration   string `json:"duration"`
			Mbid       string `json:"mbid"`
			URL        string `json:"url"`
			Streamable struct {
				Text      string `json:"#text"`
				Fulltrack string `json:"fulltrack"`
			} `json:"streamable"`
			Artist     artist  `json:"artist"`
			Image      []image `json:"image"`
			Attributes struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
		} `json:"track"`
		Attributes struct {
			Tag        string `json:"tag"`
			Page       string `json:"page"`
			PerPage    string `json:"perPage"`
			TotalPages string `json:"totalPages"`
			Total      string `json:"total"`
		} `json:"@attr"`
	} `json:"tracks"`
}

type tagWeeklyChartList struct {
	WeeklyChartList struct {
		Chart []struct {
			Text string `json:"#text"`
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"chart"`
		Attributes struct {
			Tag string `json:"tag"`
		} `json:"@attr"`
	} `json:"weeklychartlist"`
}
