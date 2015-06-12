package main

import (
	"flag"
	"github.com/cj123/robot/web"
)

var address string

func init() {
	// parse the flags
	flag.StringVar(&address, "a", "0.0.0.0:80", "the address on which to run robot's web interface")
	flag.Parse()
}

func main() {
	// start the web server
	web.Start(address)
}
