package user

import (
	"strconv"

	"git.maych.in/thunderbottom/lastfm-go"
)

// GetFriends fetches a list of friends from LastFM
func (u *User) GetFriends(page int) (fi *friendInfo, err error) {
	params := map[string]string{
		"user":  u.Username,
		"limit": u.api.GetLimit(),
		"page":  strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "user.getfriends",
		Params:   params,
		Response: &fi,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// GetInfo fetches user information from LastFM
func (u *User) GetInfo() (ui *userInfo, err error) {
	params := map[string]string{
		"user": u.Username,
	}
	p := &lastfm.Provider{
		Method:   "user.getinfo",
		Params:   params,
		Response: &ui,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// GetLovedTracks fetches tracks loved by the user from LastFM
func (u *User) GetLovedTracks(page int) (lt *lovedTracks, err error) {
	params := map[string]string{
		"user":  u.Username,
		"limit": u.api.GetLimit(),
		"page":  strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "user.getlovedtracks",
		Params:   params,
		Response: &lt,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// GetPersonalTags fetches metadata for the user's user-applied tags
// from LastFM, for the provided tag and taggingType
//
// taggingType needs to be either of `artist`, `album`, or `track`.
func (u *User) GetPersonalTags(tag string, taggingType string, page int) (pt *personalTags, err error) {
	params := map[string]string{
		"tag":         tag,
		"taggingtype": taggingType,
		"user":        u.Username,
		"page":        strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "user.getpersonaltags",
		Params:   params,
		Response: &pt,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// GetRecentTracks fetches a list of recent tracks listened to
// by the user from LastFM. Includes the current playing track.
func (u *User) GetRecentTracks(extended bool, page int) (rt *recentTracks, err error) {
	params := map[string]string{
		"user":     u.Username,
		"limit":    u.api.GetLimit(),
		"extended": u.api.Bool2strint(extended),
		"page":     strconv.Itoa(page),
	}
	p := &lastfm.Provider{
		Method:   "user.getrecenttracks",
		Params:   params,
		Response: &rt,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// GetTopAlbums fetches a list of top albums listened to by the user
// from LastFM for the specified period.
//
func (u *User) GetTopAlbums(period string, page int) (ta *topAlbums, err error) {
	params := map[string]string{
		"user":   u.Username,
		"limit":  u.api.GetLimit(),
		"page":   strconv.Itoa(page),
		"period": period,
	}
	p := &lastfm.Provider{
		Method:   "user.gettopalbums",
		Params:   params,
		Response: &ta,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// GetTopArtists fetches a list of top artists listened to by the user
// from LastFM for the specified period.
//
// period needs to be either of `overall`, `7day`, `1month`, `3month`,
// `6month`, `12month`.
func (u *User) GetTopArtists(period string, page int) (ta *topArtists, err error) {
	params := map[string]string{
		"user":   u.Username,
		"limit":  u.api.GetLimit(),
		"page":   strconv.Itoa(page),
		"period": period,
	}
	p := &lastfm.Provider{
		Method:   "user.gettopartists",
		Params:   params,
		Response: &ta,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// GetTopTracks fetches a list of top tracks listened to by the user
//from LastFM for the specified period.
//
// period needs to be either of `overall`, `7day`, `1month`, `3month`,
// `6month`, `12month`.
func (u *User) GetTopTracks(period string, page int) (tt *topTracks, err error) {
	params := map[string]string{
		"user":   u.Username,
		"limit":  u.api.GetLimit(),
		"page":   strconv.Itoa(page),
		"period": period,
	}
	p := &lastfm.Provider{
		Method:   "user.gettoptracks",
		Params:   params,
		Response: &tt,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// GetTopTags fetches top tags used by the user on LastFM.
func (u *User) GetTopTags() (tt *topTags, err error) {
	params := map[string]string{
		"user":  u.Username,
		"limit": u.api.GetLimit(),
	}
	p := &lastfm.Provider{
		Method:   "user.gettoptags",
		Params:   params,
		Response: &tt,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// GetWeeklyAlbumChart fetches album chart for the user from LastFM
// for the given date range (from, to)
//
// from and to need to be the time period for the album chart to fetch
// in unixtime.
func (u *User) GetWeeklyAlbumChart(from, to int64) (wac *weeklyAlbumChart, err error) {
	params := map[string]string{
		"user": u.Username,
		"from": strconv.FormatInt(from, 10),
		"to":   strconv.FormatInt(to, 10),
	}
	p := &lastfm.Provider{
		Method:   "user.getweeklyalbumchart",
		Params:   params,
		Response: &wac,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// GetWeeklyArtistChart fetches artist chart for the user from LastFM
// for the given date range (from, to)
//
// from and to need to be the time period for the artist chart to fetch
// in unixtime.
func (u *User) GetWeeklyArtistChart(from, to int64) (wac *weeklyArtistChart, err error) {
	params := map[string]string{
		"user": u.Username,
		"from": strconv.FormatInt(from, 10),
		"to":   strconv.FormatInt(to, 10),
	}
	p := &lastfm.Provider{
		Method:   "user.getweeklyartistchart",
		Params:   params,
		Response: &wac,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// GetWeeklyChartList fetches a list of available charts for the user
// from LastFM, expressed as date range (from, to) in unixtime.
func (u *User) GetWeeklyChartList() (wcl *weeklyChartList, err error) {
	params := map[string]string{
		"user": u.Username,
	}
	p := &lastfm.Provider{
		Method:   "user.getweeklychartlist",
		Params:   params,
		Response: &wcl,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// GetWeeklyTrackChart fetches track chart for the user from LastFM
// for the given date range (from, to)
//
// from and to need to be the time period for the track chart to fetch
// in unixtime.
func (u *User) GetWeeklyTrackChart(from, to int64) (wtc *weeklyTrackChart, err error) {
	params := map[string]string{
		"user": u.Username,
		"from": strconv.FormatInt(from, 10),
		"to":   strconv.FormatInt(to, 10),
	}
	p := &lastfm.Provider{
		Method:   "user.getweeklytrackchart",
		Params:   params,
		Response: &wtc,
		Type:     "GET",
	}
	err = u.api.Request(p)

	return
}

// New returns an instance of the `user` API endpoint functions for LastFM.
func New(client *lastfm.Client, username string) (user *User) {
	user = &User{
		api:      client,
		Username: username,
	}
	return
}
