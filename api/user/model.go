package user

import (
	"git.maych.in/thunderbottom/lastfm-go"
)

// User represents a structure to help query the `user` LastFM API functions.
type User struct {
	api      *lastfm.Client
	Username string
}

type artist struct {
	Image      []image `json:"image,omitempty"`
	Mbid       string  `json:"mbid,omitempty"`
	Name       string  `json:"name,omitempty"`
	Streamable string  `json:"streamable,omitempty"`
	Text       string  `json:"#text,omitempty"`
	URL        string  `json:"url,omitempty"`
}

type attributes struct {
	Page       string `json:"page"`
	PerPage    string `json:"perPage"`
	Tag        string `json:"tag,omitempty"`
	Total      string `json:"total"`
	TotalPages string `json:"totalPages"`
	User       string `json:"user"`
}

type friendInfo struct {
	Friends struct {
		Attributes attributes `json:"@attr"`
		User       []struct {
			Bootstrap  string     `json:"bootstrap"`
			Country    string     `json:"country"`
			Image      []image    `json:"image"`
			Name       string     `json:"name"`
			Playcount  string     `json:"playcount"`
			Playlists  string     `json:"playlists"`
			Realname   string     `json:"realname"`
			Registered registered `json:"registered"`
			Subscriber string     `json:"subscriber"`
			Type       string     `json:"type"`
			URL        string     `json:"url"`
		} `json:"user"`
	} `json:"friends"`
}

type image struct {
	Size string `json:"size"`
	Text string `json:"#text"`
}

type recentTracks struct {
	RecentTracks struct {
		Attributes attributes `json:"@attr"`
		Tracks     []track    `json:"track"`
	} `json:"recenttracks"`
}

type registered struct {
	Text     string `json:"#text"`
	Unixtime string `json:"unixtime"`
}

type track struct {
	Album struct {
		Mbid string `json:"mbid"`
		Text string `json:"#text"`
	} `json:"album"`
	Artist     artist `json:"artist"`
	Attributes struct {
		NowPlaying string `json:"nowplaying"`
	} `json:"@attr,omitempty"`
	Date struct {
		Uts  string `json:"uts"`
		Text string `json:"#text"`
	} `json:"date,omitempty"`
	Image      []image `json:"image"`
	Loved      string  `json:"loved"`
	Mbid       string  `json:"mbid"`
	Name       string  `json:"name"`
	Streamable string  `json:"streamable"`
	URL        string  `json:"url"`
}

type userInfo struct {
	User struct {
		Age        string  `json:"age"`
		Bootstrap  string  `json:"bootstrap"`
		Country    string  `json:"country"`
		Gender     string  `json:"gender"`
		Image      []image `json:"image"`
		Name       string  `json:"name"`
		Playcount  string  `json:"playcount"`
		Playlists  string  `json:"playlists"`
		Realname   string  `json:"realname"`
		Registered struct {
			Text     string `json:"#text"`
			Unixtime string `json:"#text"`
		} `json:"registered"`
		Subscriber string `json:"subscriber"`
		Type       string `json:"type"`
		URL        string `json:"url"`
	} `json:"user"`
}

type personalTags struct {
	Tags struct {
		Attributes attributes `json:"@attr"`
		Artists    struct {
			Artist []artist `json:"artist"`
		} `json:"taggings"`
		Tracks struct {
			Track []struct {
				Artist     artist  `json:"artist"`
				Duration   string  `json:"duration"`
				Image      []image `json:"image"`
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
}

type lovedTracks struct {
	LovedTracks struct {
		Attributes attributes `json:"@attr"`
		Track      []struct {
			Artist artist `json:"artist"`
			Mbid   string `json:"mbid"`
			Date   struct {
				Uts  string `json:"uts"`
				Text string `json:"#text"`
			} `json:"date"`
			URL        string  `json:"url"`
			Image      []image `json:"image"`
			Name       string  `json:"name"`
			Streamable struct {
				Fulltrack string `json:"fulltrack"`
				Text      string `json:"#text"`
			} `json:"streamable"`
		} `json:"track"`
	} `json:"lovedtracks"`
}

type topAlbums struct {
	TopAlbums struct {
		Album []struct {
			Artist struct {
				URL  string `json:"url"`
				Name string `json:"name"`
				Mbid string `json:"mbid"`
			} `json:"artist"`
			Attributes struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
			Image     []image `json:"image"`
			Playcount string  `json:"playcount"`
			URL       string  `json:"url"`
			Name      string  `json:"name"`
			Mbid      string  `json:"mbid"`
		} `json:"album"`
		Attributes attributes `json:"@attr"`
	} `json:"topalbums"`
}

type topArtists struct {
	TopArtists struct {
		Artist []struct {
			Attributes struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
			Mbid       string  `json:"mbid"`
			URL        string  `json:"url"`
			Playcount  string  `json:"playcount"`
			Image      []image `json:"image"`
			Name       string  `json:"name"`
			Streamable string  `json:"streamable"`
		} `json:"artist"`
		Attributes attributes `json:"@attr"`
	} `json:"topartists"`
}

type topTracks struct {
	TopTracks struct {
		Attributes attributes `json:"@attr"`
		Track      []struct {
			Attributes struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
			Duration   string  `json:"duration"`
			Playcount  string  `json:"playcount"`
			Artist     artist  `json:"artist"`
			Image      []image `json:"image"`
			Streamable struct {
				Fulltrack string `json:"fulltrack"`
				Text      string `json:"#text"`
			} `json:"streamable"`
			Mbid string `json:"mbid"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"track"`
	} `json:"toptracks"`
}

type topTags struct {
	TopTags struct {
		Tag []struct {
			Name  string `json:"name"`
			Count string `json:"count"`
			URL   string `json:"url"`
		} `json:"tag"`
		Attributes struct {
			User string `json:"user"`
		} `json:"@attr"`
	} `json:"toptags"`
}

type weeklyAlbumChart struct {
	WeeklyAlbumChart struct {
		Album []struct {
			Artist     artist `json:"artist"`
			Attributes struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
			Mbid      string `json:"mbid"`
			Playcount string `json:"playcount"`
			Name      string `json:"name"`
			URL       string `json:"url"`
		} `json:"album"`
		Attributes struct {
			User string `json:"user"`
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"@attr"`
	} `json:"weeklyalbumchart"`
}

type weeklyArtistChart struct {
	WeeklyArtistChart struct {
		Artist []struct {
			Attributes struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
			Mbid      string `json:"mbid"`
			Playcount string `json:"playcount"`
			Name      string `json:"name"`
			URL       string `json:"url"`
		} `json:"artist"`
		Attributes struct {
			User string `json:"user"`
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"@attr"`
	} `json:"weeklyartistchart"`
}

type weeklyChartList struct {
	WeeklyChartList struct {
		Chart []struct {
			Text string `json:"#text"`
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"chart"`
		Attributes struct {
			User string `json:"user"`
		} `json:"@attr"`
	} `json:"weeklychartlist"`
}

type weeklyTrackChart struct {
	WeeklyTrackChart struct {
		Attributes struct {
			User string `json:"user"`
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"@attr"`
		Track []struct {
			Artist     artist `json:"artist"`
			Attributes struct {
				Rank string `json:"rank"`
			} `json:"@attr"`
			Mbid      string  `json:"mbid"`
			URL       string  `json:"url"`
			Image     []image `json:"image"`
			Name      string  `json:"name"`
			Playcount string  `json:"playcount"`
		} `json:"track"`
	} `json:"weeklytrackchart"`
}
