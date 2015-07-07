package initio

// super simple (and probably broken) software pulse width modulation
// with thanks to WiringPi
// (specifically https://github.com/WiringPi/WiringPi/blob/master/wiringPi/softPwm.c)

import (
	"time"
	"github.com/stianeikeland/go-rpio"
)

// a gpio pin, with added pwm
type PWMPin struct {
	rpio.Pin
	active bool
}

const RANGE = 100

func NewPWMPin(pin int) *PWMPin{
	return &PWMPin{rpio.Pin(pin), false}
}

// software pwm on the motors, laziness rules all
func (p *PWMPin) pwm(speed uint8) {

	p.Output()

	if !p.active {
		p.active = true
	}

	go func() {
		for p.active {

			p.High()

			time.Sleep(time.Duration(speed * 100) * time.Microsecond)

			p.Pin.Low()

			time.Sleep(time.Duration((RANGE - speed) * 100) * time.Microsecond)
		}
	}()
}

func (p *PWMPin) stop() {
	p.active = false
}

// override the pin's usual low, to stop pwm
func (p *PWMPin) Low() {
	p.stop()
	p.Pin.Low()
}
