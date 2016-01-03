// +build !linux,!arm

package initio

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// this will be an exact mirror of initio externally, but
// will use the web API running on the actual robot, allowing
// for more rapid remote development
func init() {
	// catch ^C, and cleanup appropriately
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for _ = range c {
			// sig is a ^C, handle it
			log.Println("^C detected, closing cleanly...")

			Cleanup()

			os.Exit(0)
		}
	}()
}

var baseURL = "http://192.168.79.21/"

func SetBaseURL(url string) {
	baseURL = url
}

// make the request to the URL, returning success or error if failed
func makeRequest(url string) ([]byte, int, error) {
	log.Println("making request to:", (baseURL + url))

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
	// we can open a new instance of motor - they're the same pins
	m := NewMotor()
	m.Stop()     // stop motors
	StopServos() // stop servos
	//	rpio.Close()
}
