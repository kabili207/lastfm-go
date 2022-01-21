// Package lastfm implements the LastFM API 2.0
//
// This package is a wrapper for the LastFM API: https://www.last.fm/api
//
// The wrapper is built to be modular, and each LastFM API endpoint method
// is a separate package. Following API methods are covered by the package:
//
// album: https://godoc.org/git.maych.in/thunderbottom/lastfm-go/api/album
//
// artist: https://godoc.org/git.maych.in/thunderbottom/lastfm-go/api/artist
//
// chart: https://godoc.org/git.maych.in/thunderbottom/lastfm-go/api/chart
//
// geo: https://godoc.org/git.maych.in/thunderbottom/lastfm-go/api/geo
//
// library: https://godoc.org/git.maych.in/thunderbottom/lastfm-go/api/library
//
// tag: https://godoc.org/git.maych.in/thunderbottom/lastfm-go/api/tag
//
// track: https://godoc.org/git.maych.in/thunderbottom/lastfm-go/api/track
//
// user: https://godoc.org/git.maych.in/thunderbottom/lastfm-go/api/user
package lastfm

import (
	"net/http"
	"strconv"
	"time"
)

const (
	apiBaseURL = "https://ws.audioscrobbler.com/2.0/"
	authURL    = "https://www.last.fm/api/auth/"
)

// New returns an instance of the LastFM Client.
// The instance also includes an HTTP client for querying the LastFM API, with timeout set to 10 seconds.
// The page limit for requests is set to 50 by default, which can be changed using SetLimit
func New(apiKey, apiSecret string) (client Client) {
	client = Client{
		APIKey:    apiKey,
		APISecret: apiSecret,
		limit:     50,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	return
}

// Request uses an instance of Client to perform GET and POST requests on the LastFM API endpoints.
// provider must be an instance of Provider containing the API Method, parameters for the API call,
// and the request Type. Optionally, the provider should also include an interface to Unmarshal the
// request response.
//
// This function is usually called from functions abstracting the LastFM API.
func (client *Client) Request(provider *Provider) (err error) {
	req, err := http.NewRequest(provider.Type, apiBaseURL, nil)
	if err != nil {
		return err
	}

	params := req.URL.Query()
	params.Add("method", provider.Method)
	params.Add("api_key", client.APIKey)
	for key, value := range provider.Params {
		params.Add(key, value)
	}

	switch provider.Type {
	case "GET":
		params.Add("format", "json")
	case "POST":
		if client.sessionKey != "" {
			params.Add("sk", client.sessionKey)
		}
		signature := client.generateSignature(params)
		params.Add("api_sig", signature)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	if client.useragent != "" {
		req.Header.Set("User-Agent", client.useragent)
	}
	req.URL.RawQuery = params.Encode()
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return client.parseError(resp, provider)
	}

	if provider.Response != nil {
		err = client.parseResponse(resp, provider.Response)
	}

	return
}

// SetUserAgent sets the User-Agent header to be used while querying the LastFM API.
func (client *Client) SetUserAgent(useragent string) {
	client.useragent = useragent
}

// GetUserAgent returns the currently set User-Agent header
func (client *Client) GetUserAgent() string {
	return client.useragent
}

// SetLimit sets the page limit on the Client for LastFM API requests. The limit defaults to 50 for
// new Client instances or when set to <= 0.
func (client *Client) SetLimit(limit int) {
	if limit <= 0 {
		limit = 50
	}
	client.limit = limit
}

// GetLimit returns a string representation of the current page limit set on the Client.
// This function is usually called from functions abstracting the LastFM API where page limits are required.
func (client *Client) GetLimit() (limit string) {
	return strconv.Itoa(client.limit)
}
