package artist

import (
	"git.maych.in/thunderbottom/lastfm-go"
)

// Artist represents a structure to help query the `artist` LastFM API functions.
type Artist struct {
	api         *lastfm.Client
	autocorrect string
	Username    string
}

type artist struct {
	Name       string  `json:"name"`
	Image      []image `json:"image,omitempty"`
	Listeners  string  `json:"listeners,omitempty"`
	Match      string  `json:"match,omitempty"`
	Mbid       string  `json:"mbid,omitempty"`
	Streamable string  `json:"streamable,omitempty"`
	URL        string  `json:"url"`
}

type image struct {
	Text string `json:"#text"`
	Size string `json:"size"`
}

type tags struct {
	Count int    `json:"count,omitempty"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type artisttag struct {
	Tag        []tags `json:"tag"`
	Attributes struct {
		Artist string `json:"artist"`
	} `json:"@attr"`
}

type artistCorrection struct {
	Corrections struct {
		Correction struct {
			Artist     artist `json:"artist"`
			Attributes struct {
				Index string `json:"index"`
			} `json:"@attr"`
		} `json:"correction"`
	} `json:"corrections"`
}

type artistInfo struct {
	Artist struct {
		Name       string  `json:"name"`
		Mbid       string  `json:"mbid"`
		URL        string  `json:"url"`
		Image      []image `json:"image"`
		Streamable string  `json:"streamable"`
		Ontour     string  `json:"ontour"`
		Stats      struct {
			Listeners string `json:"listeners"`
			Playcount string `json:"playcount"`
		} `json:"stats"`
		Similar struct {
			Artist []artist `json:"artist"`
		} `json:"similar"`
		Tags struct {
			Tag []tags `json:"tag"`
		} `json:"tags"`
		Bio struct {
			Links struct {
				Link struct {
					Text string `json:"#text"`
					Rel  string `json:"rel"`
					Href string `json:"href"`
				} `json:"link"`
			} `json:"links"`
			Published string `json:"published"`
			Summary   string `json:"summary"`
			Content   string `json:"content"`
		} `json:"bio"`
	} `json:"artist"`
}

type artistSimilar struct {
	SimilarArtists struct {
		Artist []struct {
			Name       string  `json:"name"`
			Mbid       string  `json:"mbid"`
			Match      string  `json:"match"`
			URL        string  `json:"url"`
			Image      []image `json:"image"`
			Streamable string  `json:"streamable"`
		} `json:"artist"`
		Attributes struct {
			Artist string `json:"artist"`
		} `json:"@attr"`
	} `json:"similarartists"`
}

type artistTags struct {
	Tags artisttag `json:"tags"`
}

type artistTopAlbums struct {
	TopAlbums struct {
		Album []struct {
			Name      string  `json:"name"`
			Playcount int     `json:"playcount"`
			Mbid      string  `json:"mbid,omitempty"`
			URL       string  `json:"url"`
			Artist    artist  `json:"artist"`
			Image     []image `json:"image"`
		} `json:"album"`
		Attributes struct {
			Artist     string `json:"artist"`
			Page       string `json:"page"`
			PerPage    string `json:"perPage"`
			TotalPages string `json:"totalPages"`
			Total      string `json:"total"`
		} `json:"@attr"`
	} `json:"topalbums"`
}

type artistTopTags struct {
	TopTags artisttag `json:"toptags"`
}

type artistTopTracks struct {
	TopTracks struct {
		Track []struct {
			Name       string  `json:"name"`
			Playcount  string  `json:"playcount"`
			Listeners  string  `json:"listeners"`
			Mbid       string  `json:"mbid,omitempty"`
			URL        string  `json:"url"`
			Streamable string  `json:"streamable"`
			Artist     artist  `json:"artist"`
			Image      []image `json:"image"`
			Attributes struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
		} `json:"track"`
		Attributes struct {
			Artist     string `json:"artist"`
			Page       string `json:"page"`
			PerPage    string `json:"perPage"`
			TotalPages string `json:"totalPages"`
			Total      string `json:"total"`
		} `json:"@attr"`
	} `json:"toptracks"`
}

type artistSearch struct {
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
		Artistmatches          struct {
			Artist []artist `json:"artist"`
		} `json:"artistmatches"`
		Attributes struct {
			For string `json:"for"`
		} `json:"@attr"`
	} `json:"results"`
}
