package lastfm

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"unsafe"
)

// Bool2strint returns a string representation of the integer value
// for the provided boolean input.
func (client *Client) Bool2strint(b bool) string {
	return strconv.Itoa(int(uint8(*(*uint8)(unsafe.Pointer(&b)))))
}

func (client *Client) parseResponse(resp *http.Response, response interface{}) (err error) {
	read, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	switch resp.Header.Get("Content-Type") {
	case "application/json":
		err = json.Unmarshal(read, &response)
	default:
		base := xmlBase{}
		err = xml.Unmarshal(read, &base)
		if err != nil {
			return
		}
		if len(base.Inner) > 0 {
			err = xml.Unmarshal(base.Inner, &response)
		}
	}
	return
}

func (client *Client) parseError(resp *http.Response, provider *Provider) (err error) {
	respErr := &Error{}
	err = client.parseResponse(resp, respErr)
	if err != nil {
		return
	}
	switch respErr.ErrorCode {
	case 2, 3, 5, 6, 7, 13:
		return fmt.Errorf("%v: %v", provider.Method, respErr.Message)
	case 10, 26:
		return fmt.Errorf("%v: %v", client.APIKey, respErr.Message)
	default:
		return fmt.Errorf("Error Code: %v\nMessage:%v", respErr.ErrorCode, respErr.Message)
	}
}
