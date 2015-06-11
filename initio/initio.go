package initio

import (
	"github.com/stianeikeland/go-rpio"
)

func init() {
	err := rpio.Open()

	if err != nil {
		panic(err)
	}
}

func Cleanup() {
	rpio.Close()
}
