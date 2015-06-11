package initio

import (
	"github.com/stianeikeland/go-rpio"
)

func init() {
	err := rpio.Open()

	if err != nil {
		panic(err)
	}

	StartServos()

}

func Cleanup() {
	//Stop()       // stop motors
	StopServos() // stop servos
	rpio.Close()
}
