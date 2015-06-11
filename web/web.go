package web

// provides a way of controlling/reading the sensors over the web

import (
	"encoding/json"
	"fmt"
	"github.com/cj123/robot/initio" // our robot commands
	"net/http"
	"strings"
)

func Start() {
	http.HandleFunc("/api/", apihandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "static/"+r.URL.Path[1:]) }) // static files
	http.ListenAndServe(":8080", nil)
}

var routes = map[string]func(s []string, w http.ResponseWriter, r *http.Request){
	"sonar":  sonar,
	"ir":     ir,
	"motors": motors,
	"servos": servos,
}

func apihandler(w http.ResponseWriter, r *http.Request) {

	// get an array of parts of URL, as split by /
	urlParts := strings.Split(r.URL.Path[len("/api/"):], "/")

	if fn, ok := routes[strings.ToLower(urlParts[0])]; ok {
		fn(urlParts, w, r)
	}
}

// sonar data
func sonar(s []string, w http.ResponseWriter, r *http.Request) {
	if len(s) > 1 && s[1] == "distance" {
		fmt.Fprintf(w, "%d\n", initio.GetDistance())
	}
}

// JSON for the IR sensors
type IRJSON struct {
	Left      bool
	Right     bool
	LeftLine  bool
	RightLine bool
}

// ir data
func ir(s []string, w http.ResponseWriter, r *http.Request) {
	if len(s) > 1 && s[1] == "distance" {
		fmt.Fprintf(w, "%d\n", initio.GetDistance())
	}
}

// motors data
func motors(s []string, w http.ResponseWriter, r *http.Request) {

}

// servos data
func servos(s []string, w http.ResponseWriter, r *http.Request) {

}
