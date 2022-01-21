package album

import (
	"git.maych.in/thunderbottom/lastfm-go"
)

// Album represents a structure to help query the `album` LastFM API functions.
type Album struct {
	api         *lastfm.Client
	autocorrect string
	Username    string
}

type artist struct {
	Name string `json:"name"`
	Mbid string `json:"mbid"`
	URL  string `json:"url"`
}

type attributes struct {
	Artist string `json:"artist"`
	Album  string `json:"album"`
}

type image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

type streamable struct {
	Text      string `json:"#text"`
	Fulltrack string `json:"fulltrack"`
}

type tags struct {
	Count int    `json:"count,omitempty"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type albumInfo struct {
	Album struct {
		Name          string  `json:"name"`
		Artist        string  `json:"artist"`
		URL           string  `json:"url"`
		Image         []image `json:"image"`
		Listeners     string  `json:"listeners"`
		Playcount     string  `json:"playcount"`
		Userplaycount string  `json:"userplaycount"`
		Tracks        struct {
			Track []struct {
				Artist     artist     `json:"artist"`
				Name       string     `json:"name"`
				URL        string     `json:"url"`
				Duration   string     `json:"duration"`
				Streamable streamable `json:"streamable"`
				Attributes struct {
					Rank string `json:"rank"`
				} `json:"@attr"`
			} `json:"track"`
		} `json:"tracks"`
		Tags struct {
			Tag []tags `json:"tag"`
		} `json:"tags"`
		Wiki struct {
			Published string `json:"published"`
			Summary   string `json:"summary"`
			Content   string `json:"content"`
		} `json:"wiki"`
	} `json:"album"`
}

type albumTags struct {
	Tags struct {
		Tag        []tags     `json:"tag"`
		Attributes attributes `json:"@attr"`
	} `json:"tags"`
}

type albumTopTags struct {
	TopTags struct {
		Tag        []tags     `json:"tag"`
		Attributes attributes `json:"@attr"`
	} `json:"toptags"`
}

type albumSearch struct {
	Results struct {
		OpensearchQuery struct {
			Text        string `json:"#text"`
			Role        string `json:"role"`
			SearchTerms string `json:"searchTerms"`
			StartPage   string `json:"startPage"`
		} `json:"opensearch:Query"`
		OpensearchTotalResults string `json:"opensearch:totalResults"`
		OpensearchStartIndex   string `json:"opensearch:startIndex"`
		OpensearchItemsPerPage string `json:"opensearch:itemsPerPage"`
		Albummatches           struct {
			Album []struct {
				Name       string  `json:"name"`
				Artist     string  `json:"artist"`
				URL        string  `json:"url"`
				Image      []image `json:"image"`
				Streamable string  `json:"streamable"`
				Mbid       string  `json:"mbid"`
			} `json:"album"`
		} `json:"albummatches"`
		Attr struct {
			For string `json:"for"`
		} `json:"@attr"`
	} `json:"results"`
}
