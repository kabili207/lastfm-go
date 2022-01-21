package lastfm

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"sort"
)

func (client *Client) generateSignature(params url.Values) (signature string) {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var sigParams string
	for _, key := range keys {
		sigParams += key + params.Get(key)
	}
	sigParams += client.APISecret

	md5Hash := md5.New()
	md5Hash.Write([]byte(sigParams))
	signature = hex.EncodeToString(md5Hash.Sum(nil))
	return
}

// Login creates a web service session for the LastFM user by authenticating
// using LastFM login credentials, and sets the session key within the LastFM Client.
//
// username and password must be a valid LastFM user credentials.
//
// This function must be called once before executing any function performing POST request
// on the LastFM API.
func (client *Client) Login(username string, password string) (err error) {
	client.sessionKey = ""
	params := map[string]string{
		"username": username,
		"password": password,
	}
	var auth Auth
	p := &Provider{
		Method:   "auth.getmobilesession",
		Params:   params,
		Response: &auth,
		Type:     "POST",
	}
	err = client.Request(p)
	client.sessionKey = auth.Key
	return
}

// Logout clears the current web service session and logs the user out of LastFM
func (client *Client) Logout() {
	client.sessionKey = ""
}
