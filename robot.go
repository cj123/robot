package main

import (
	"flag"
	"fmt"
	"github.com/cj123/robot/collisions"
	"github.com/cj123/robot/web"
)

var address string

func init() {
	// parse the flags
	flag.StringVar(&address, "a", "0.0.0.0:80", "the address on which to run robot's web interface")

	flag.Parse()

	fmt.Println("Robot initialised")
}

func main() {

	c := make(chan bool, 1)

	go func() {
		c <- web.Start(address)
	}()

	collisions.DoFunkyCollisionAvoidance()

}
