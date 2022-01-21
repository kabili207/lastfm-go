package track

import (
	"encoding/xml"

	"git.maych.in/thunderbottom/lastfm-go"
)

// Track represents a structure to help query the `track` LastFM API functions.
type Track struct {
	api         *lastfm.Client
	autocorrect string
	Username    string
}

type artist struct {
	Name string `json:"name"`
	Mbid string `json:"mbid"`
	URL  string `json:"url"`
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

type tracktag struct {
	Tag        []tags `json:"tag"`
	Attributes struct {
		Artist string `json:"artist"`
		Track  string `json:"track"`
	} `json:"@attr"`
}

type trackCorrection struct {
	Corrections struct {
		Correction struct {
			Track struct {
				Name   string `json:"name"`
				URL    string `json:"url"`
				Artist artist `json:"artist"`
			} `json:"track"`
			Attributes struct {
				Index           string `json:"index"`
				Artistcorrected string `json:"artistcorrected"`
				Trackcorrected  string `json:"trackcorrected"`
			} `json:"@attr"`
		} `json:"correction"`
	} `json:"corrections"`
}

type trackInfo struct {
	Track struct {
		Name       string     `json:"name"`
		Mbid       string     `json:"mbid"`
		URL        string     `json:"url"`
		Duration   string     `json:"duration"`
		Streamable streamable `json:"streamable"`
		Listeners  string     `json:"listeners"`
		Playcount  string     `json:"playcount"`
		Artist     artist     `json:"artist"`
		Album      struct {
			Artist     string  `json:"artist"`
			Title      string  `json:"title"`
			Mbid       string  `json:"mbid"`
			URL        string  `json:"url"`
			Image      []image `json:"image"`
			Attributes struct {
				Position string `json:"position"`
			} `json:"@attr"`
		} `json:"album"`
		Userplaycount string `json:"userplaycount"`
		Userloved     string `json:"userloved"`
		Toptags       struct {
			Tag []tags `json:"tag"`
		} `json:"toptags"`
		Wiki struct {
			Published string `json:"published"`
			Summary   string `json:"summary"`
			Content   string `json:"content"`
		} `json:"wiki"`
	} `json:"track"`
}

type trackSimilar struct {
	Similartracks struct {
		Track []struct {
			Name       string     `json:"name"`
			Playcount  int        `json:"playcount"`
			Mbid       string     `json:"mbid,omitempty"`
			Match      float64    `json:"match"`
			URL        string     `json:"url"`
			Streamable streamable `json:"streamable"`
			Duration   int        `json:"duration,omitempty"`
			Artist     artist     `json:"artist"`
		} `json:"track"`
		Attributes struct {
			Artist string `json:"artist"`
		} `json:"@attr"`
	} `json:"similartracks"`
}

type trackTags struct {
	Tags tracktag `json:"tags"`
}

type trackTopTags struct {
	Tags tracktag `json:"toptags"`
}

type trackScrobble struct {
	XMLName   xml.Name `xml:"scrobbles"`
	Accepted  string   `xml:"accepted,attr"`
	Ignored   string   `xml:"ignored,attr"`
	Scrobbles []struct {
		Track struct {
			Corrected string `xml:"corrected,attr"`
			Name      string `xml:",chardata"`
		} `xml:"track"`
		Artist struct {
			Corrected string `xml:"corrected,attr"`
			Name      string `xml:",chardata"`
		} `xml:"artist"`
		Album struct {
			Corrected string `xml:"corrected,attr"`
			Name      string `xml:",chardata"`
		} `xml:"album"`
		AlbumArtist struct {
			Corrected string `xml:"corrected,attr"`
			Name      string `xml:",chardata"`
		} `xml:"albumArtist"`
		TimeStamp      string `xml:"timestamp"`
		IgnoredMessage struct {
			Corrected string `xml:"corrected,attr"`
			Body      string `xml:",chardata"`
		} `xml:"ignoredMessage"`
	} `xml:"scrobble"`
}

type trackSearch struct {
	Results struct {
		OpensearchQuery struct {
			Text      string `json:"#text"`
			Role      string `json:"role"`
			StartPage string `json:"startPage"`
		} `json:"opensearch:Query"`
		OpensearchTotalResults string `json:"opensearch:totalResults"`
		OpensearchStartIndex   string `json:"opensearch:startIndex"`
		OpensearchItemsPerPage string `json:"opensearch:itemsPerPage"`
		Trackmatches           struct {
			Track []struct {
				Name       string  `json:"name"`
				Artist     string  `json:"artist"`
				URL        string  `json:"url"`
				Streamable string  `json:"streamable"`
				Listeners  string  `json:"listeners"`
				Image      []image `json:"image"`
				Mbid       string  `json:"mbid"`
			} `json:"track"`
		} `json:"trackmatches"`
		Attributes struct {
		} `json:"@attr"`
	} `json:"results"`
}

type trackUpdateNowPlaying struct {
	XMLName xml.Name `xml:"nowplaying"`
	Track   struct {
		Corrected string `xml:"corrected,attr"`
		Name      string `xml:",chardata"`
	} `xml:"track"`
	Artist struct {
		Corrected string `xml:"corrected,attr"`
		Name      string `xml:",chardata"`
	} `xml:"artist"`
	Album struct {
		Corrected string `xml:"corrected,attr"`
		Name      string `xml:",chardata"`
	} `xml:"album"`
	AlbumArtist struct {
		Corrected string `xml:"corrected,attr"`
		Name      string `xml:",chardata"`
	} `xml:"albumArtist"`
	IgnoredMessage struct {
		Corrected string `xml:"corrected,attr"`
		Body      string `xml:",chardata"`
	} `xml:"ignoredMessage"`
}
