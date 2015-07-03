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

	// initialise the two
	pan := NewServo(Pan)
	tilt := NewServo(Tilt)

	// set them to known values
	pan.Set(DEFAULT_VAL)
	tilt.Set(DEFAULT_VAL)

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
	// we can open a new instance of motor - they're the same pins
	m := NewMotor()
	m.Stop()     // stop motors
	StopServos() // stop servos
	//	rpio.Close()
}
