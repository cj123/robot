package web

// provides a way of controlling/reading the sensors over the web

import (
	"encoding/json"
	"fmt"
	"github.com/cj123/robot/initio" // our robot commands
	"net/http"
	"strconv"
	"strings"
)

func Start(address string) bool {
	http.HandleFunc("/api/", apihandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path[1:])
		http.ServeFile(w, r, "web/static/"+r.URL.Path[1:])
	}) // static files
	http.ListenAndServe(address, nil)

	return true
}

var routes = map[string]func(s []string, w http.ResponseWriter, r *http.Request){
	"sonar":     sonar,
	"ir":        ir,
	"motors":    motors,
	"servos":    servos,
	"collision": collision,
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
	Left      bool `json:"left"`
	Right     bool `json:"right"`
	LeftLine  bool `json:"leftline"`
	RightLine bool `json:"rightline"`
}

// ir data
func ir(s []string, w http.ResponseWriter, r *http.Request) {
	// instantiate the ir sensor
	irSensor := initio.NewIR()

	irs := &IRJSON{
		Left:      irSensor.Left(),
		Right:     irSensor.Right(),
		LeftLine:  irSensor.LeftLine(),
		RightLine: irSensor.RightLine(),
	}

	resp, err := json.MarshalIndent(irs, "", "  ")

	if err != nil {
		panic(err) // for now
	}

	fmt.Fprintf(w, "%s", resp)
}

// motors data
func motors(s []string, w http.ResponseWriter, r *http.Request) {
	if len(s) < 2 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	if s[1] == "forwards" {
		initio.Forward(0) // speed TODO
	} else if s[1] == "reverse" {
		initio.Reverse(0)
	} else if s[1] == "left" {
		initio.SpinLeft(0)
	} else if s[1] == "right" {
		initio.SpinRight(0)
	} else if s[1] == "stop" {
		initio.Stop()
	}
}

// servos data
func servos(s []string, w http.ResponseWriter, r *http.Request) {
	if len(s) < 3 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	servo := -1

	// which servo are we using?
	if s[1] == "pan" {
		servo = initio.Pan
	} else if s[1] == "tilt" {
		servo = initio.Tilt
	} else {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	// process the request
	if s[2] == "set" {
		val, _ := strconv.Atoi(s[3])

		// set the servo to this pos
		initio.SetServo(servo, val)
	} else if s[2] == "inc" {
		val, _ := strconv.Atoi(s[3])

		// increment the servo
		initio.IncServo(servo, val)
	} else if s[2] == "get" {
		// return servo pos
		fmt.Fprintf(w, "%d", initio.GetServo(servo))
	} else if s[2] == "reset" {
		initio.ResetServo(servo)
	}
}

func collision(s []string, w http.ResponseWriter, r *http.Request) {
	if len(s) < 2 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	if s[1] == "on" {
		// start the avoidance system

	}

	if s[2] == "off" {
		// stop the avoidance system

	}

}
