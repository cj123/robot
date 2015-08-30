// +build !linux,!arm

package initio

import (
	"io/ioutil"
	//"log"
	"net/http"
)

// this will be an exact mirror of initio externally, but
// will use the web API running on the actual robot, allowing
// for more rapid remote development
func init() {}

var baseURL = "http://192.168.79.21/"

func SetBaseURL(url string) {
	baseURL = url
}

// make the request to the URL, returning success or error if failed
func makeRequest(url string) ([]byte, int, error) {
	//log.Println("making request to:", (baseURL + url))

	resp, err := http.Get(baseURL + url)

	if err != nil {
		return nil, -1, err
	}

	document, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, -1, err
	}

	return document, resp.StatusCode, err
}

func Cleanup() {
	// do nothing. In the actual library this will do something necessary
	// to ensure proper shutdown of the robot. but here, we can assume that is
	// handled from the API
}
