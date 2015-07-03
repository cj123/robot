package web

// provides a way of controlling/reading the sensors over the web

import (
	"encoding/json"
	"fmt"
	"github.com/cj123/robot/initio" // our robot commands
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	collisionAvoidance  *bool
	motor               *initio.Motors
	sonarSensor         *initio.Sonar
	panServo, tiltServo *initio.Servo
)

func Start(address string, runCollisionAvoidance *bool) bool {
	motor = initio.NewMotor()
	sonarSensor = initio.NewSonar()

	panServo = initio.NewServo(initio.Pan)
	tiltServo = initio.NewServo(initio.Tilt)

	collisionAvoidance = runCollisionAvoidance

	log.Println("Started webserver on " + address)

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
	"domain":    getDomain,
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
		fmt.Fprintf(w, "%d\n", sonarSensor.GetDistance())
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
		motor.Forward(0) // speed TODO
	} else if s[1] == "reverse" {
		motor.Reverse(0)
	} else if s[1] == "left" {
		motor.SpinLeft(0)
	} else if s[1] == "right" {
		motor.SpinRight(0)
	} else if s[1] == "stop" {
		motor.Stop()
	}
}

// servos data
func servos(s []string, w http.ResponseWriter, r *http.Request) {
	if len(s) < 3 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	var servo *initio.Servo

	// which servo are we using?
	if s[1] == "pan" {
		servo = panServo
	} else if s[1] == "tilt" {
		servo = tiltServo
	} else {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	// process the request
	if s[2] == "set" {
		val, _ := strconv.Atoi(s[3])

		// set the servo to this pos
		servo.Set(val)
	} else if s[2] == "inc" {
		val, _ := strconv.Atoi(s[3])

		// increment the servo
		servo.Inc(val)
	} else if s[2] == "get" {
		// return servo pos
		fmt.Fprintf(w, "%d", servo.Get())
	} else if s[2] == "reset" {
		servo.Reset()
	}
}

func collision(s []string, w http.ResponseWriter, r *http.Request) {
	if len(s) < 2 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	if s[1] == "on" {
		// start the avoidance system
		*collisionAvoidance = true
	} else if s[1] == "off" {
		// stop the avoidance system
		*collisionAvoidance = false
	} else if s[1] == "get" {
		fmt.Fprintf(w, "%t", *collisionAvoidance)
	} else if s[1] == "toggle" {
		*collisionAvoidance = !*collisionAvoidance
	} else {
		http.Error(w, "Invalid Request: "+s[1], http.StatusBadRequest)
	}
}

// get the current domain
func getDomain(s []string, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.Host)
}
