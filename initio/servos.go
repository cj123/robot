// +build !linux,!arm

package initio

import (
	"fmt"
	"strconv"
)

type Servo struct {
	Pin     int // i.e. tilt or pan
	Current int // the current position
}

func NewServo(pin int) *Servo {
	return &Servo{Pin: pin}
}

// set a servo to a certain angle
func (s *Servo) Set(degrees int) {
	makeRequest(fmt.Sprintf("/api/servos/%s/set/%d", servoNames[s.Pin], degrees))
}

// get the current value the servo is at
func (s *Servo) Get() int {
	result, _, _ := makeRequest(fmt.Sprintf("/api/servos/%s/get", servoNames[s.Pin]))
	i, _ := strconv.Atoi(string(result))
	return i
}

func (s *Servo) Reset() {
	makeRequest(fmt.Sprintf("/api/servos/%s/reset", servoNames[s.Pin]))
}

// increment (or decrement) a servo by a value
func (s *Servo) Inc(increment int) {
	// TODO: stubbed
}

// stop the servos
func StopServos() {
	// do nothing
}

// start the servos
func StartServos() {
	// do nothing
}
