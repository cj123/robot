package initio

import (
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
)

const (
	Pan         = 0
	Tilt        = 1
	DEFAULT_VAL = 30

	// the maximum range of the servo
	MAX_RANGE = 180
	MIN_VAL   = -90
	MAX_VAL   = 90
)

var servoPins = [2]int{24, 25}

type Servo struct {
	Pin     int // i.e. tilt or pan
	Current int // the current position
	Active  bool
}

// create a new servo, either Pan or Tilt
func NewServo(pin int) *Servo {
	s := &Servo{Pin: pin, Active: false}

	return s
}

var fltMillisecond = float64(time.Millisecond)

// set a servo to a certain angle
// this unfortunately will not be 100% accurate due to software constraints
// but it's better than loading a huge library for the purposes of a simple bit
// of servo control.
func (s *Servo) Set(degrees int) {

	if s.Active {
		s.Active = false
	}

	// check value range
	if degrees < MIN_VAL || degrees > MAX_VAL {
		return
	}

	// open the pin as an output
	pin := rpio.Pin(servoPins[s.Pin])
	pin.Output()

	pulseWidth := float64(degrees+MAX_VAL) / float64(MAX_RANGE)
	duration := time.Duration(((fltMillisecond - (pulseWidth * fltMillisecond)) * 2) + (fltMillisecond / 2))

	log.Println("Pulse width:", pulseWidth, duration)

	go func() {

		s.Active = true

		for i := 0; i < 10 && s.Active; i++ {
			// send a pulse minimum 0.5ms, but maximum 2.5ms
			pin.High()
			timer := time.NewTimer(duration)
			<-timer.C
			pin.Low()

			// update current value
			s.Current = degrees

			// repeat every 20ms
			timer = time.NewTimer(20 * time.Millisecond)
			<-timer.C
		}

	}()
}

// get the current value the servo is at
func (s *Servo) Get() int {
	return s.Current
}

// set the servo back to it's default position
func (s *Servo) Reset() {
	s.Set(DEFAULT_VAL)
}

// increment (or decrement) a servo by a value
func (s *Servo) Inc(increment int) {
	val := s.Get()
	val += increment

	s.Set(val)
}
