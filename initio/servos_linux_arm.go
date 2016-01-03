package initio

import (
	"fmt"
	"io/ioutil"
	"os/exec" // sorry :(
)

var (
	// whether the servos are active
	servosActive = false
)

// Servo is a servo pin
type Servo struct {
	Pin     int // i.e. tilt or pan
	Current int // the current position
}

// NewServo instantiates servo on a pin
func NewServo(pin int) *Servo {
	return &Servo{Pin: pin}
}

// Set a servo to a certain angle
func (s *Servo) Set(degrees int) {
	if !servosActive {
		StartServos()
	}

	pinServod(s.Pin, degrees)
	s.Current = degrees
}

// Get the current value the servo is at
func (s *Servo) Get() int {
	return s.Current
}

// Reset the servo pin
func (s *Servo) Reset() {
	s.Set(DEFAULT_VAL)
}

// Inc increments (or decrements) a servo by a value
func (s *Servo) Inc(increment int) {
	val := s.Get()
	val += increment

	s.Set(val)
}

// StopServos stops the servos
func StopServos() {
	stopServod()
}

// StartServos starts up the servo daemon
func StartServos() {
	err := startServod()

	if err != nil {
		panic(err)
	}
}

// start servod, the servo daemon
func startServod() error {
	if servosActive {
		// already open
		return nil
	}

	// run the command ./servod --pcm --idle-timeout=20000 --p1pins="18,22"
	cmd := exec.Command("sudo", "sh", "-c", "servod --pcm --idle-timeout=20000 --p1pins=\"18,22\"")
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
		return err
	}

	servosActive = true

	return err
}

// stop servod, the servo daemon
func stopServod() error {
	if !servosActive {
		// don't close something that isn't there
		return nil
	}

	cmd := exec.Command("sudo", "sh", "-c", "killall servod")
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
		return err
	}

	servosActive = false

	return err
}

// apply servo change to the pin
func pinServod(pin int, degrees int) {
	pinString := fmt.Sprintf("%d=%d\n", pin, 50+((90-degrees)*200/180))
	//fmt.Println(pinString)
	err := ioutil.WriteFile("/dev/servoblaster", []byte(pinString), 0644)

	if err != nil {
		panic(err)
	}
}
