package main

import (
	"bytes"
	"flag"
	"github.com/cj123/robot/collisions"
	"github.com/cj123/robot/initio"
	"github.com/cj123/robot/web"
	"io"
	"log"
	"os"
)

var (
	// address of web interface
	address string

	// remote API address
	remoteAddr string

	// whether we should run collision avoidance
	runCollisionAvoidance bool

	// output for logging
	buf *bytes.Buffer
)

func init() {
	// parse the flags
	flag.StringVar(&address, "a", "0.0.0.0:80", "the address on which to run robot's web interface")
	flag.StringVar(&remoteAddr, "r", "127.0.0.1:80", "the remote address on which the robot API is running")
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

	initio.SetBaseURL(remoteAddr)

	// start the collisions model
	collisions.Start(&runCollisionAvoidance)
}
