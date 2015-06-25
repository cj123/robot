package initio

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"os"
	"os/signal"
)

func init() {
	err := rpio.Open()

	if err != nil {
		panic(err)
	}

	// setup the servos
	StartServos()

	// set them to known values
	SetServo(Tilt, DEFAULT_VAL)
	SetServo(Pan, DEFAULT_VAL)

	// catch ^C, and cleanup appropriately
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for _ = range c {
			// sig is a ^C, handle it
			log.Println("^C detected, closing cleanly...")

			Cleanup()
			rpio.Close()

			os.Exit(0)
		}
	}()
}

func Cleanup() {
	Stop()       // stop motors
	StopServos() // stop servos
	//	rpio.Close()
}
