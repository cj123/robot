package initio

import (
	"github.com/stianeikeland/go-rpio"
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
}

func Cleanup() {
	Stop()       // stop motors
	StopServos() // stop servos
	//	rpio.Close()
}
