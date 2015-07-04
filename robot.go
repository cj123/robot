package main

import (
	"flag"
	"github.com/cj123/robot/collisions"
	"github.com/cj123/robot/web"
	"log"
	"bytes"
	"os"
	"io"
)

// address of web interface
var address string

// whether we should run collision avoidance
var runCollisionAvoidance bool

var buf *bytes.Buffer

func init() {
	// parse the flags
	flag.StringVar(&address, "a", "0.0.0.0:80", "the address on which to run robot's web interface")
	flag.BoolVar(&runCollisionAvoidance, "c", false, "should we try avoid collisions?")
	flag.Parse()

	buf = new(bytes.Buffer)

	log.SetOutput(io.MultiWriter(buf, os.Stdout))

	log.Println("Robot initialised.")
	log.Printf("Collision avoidance is set to: %t\n", runCollisionAvoidance)
}

func main() {
	c := make(chan bool, 1)

	go func() {
		c <- web.Start(address, &runCollisionAvoidance, buf)
	}()

	// start the collisions model
	collisions.Start(&runCollisionAvoidance)
}
