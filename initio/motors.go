package initio

import (
	"fmt"

	"github.com/stianeikeland/go-rpio"
)

const (
	RightMotor1 = 8 // right motor pins 24
	RightMotor2 = 7 // 26

	LeftMotor1 = 10 // left motor pins: 19
	LeftMotor2 = 9  // 21
)

type Motors struct {
	// the motor points
	a, b, p, q rpio.Pin
}

func NewMotor() *Motors {
	m := Motors{}

	// initialise the pins
	m.p = rpio.Pin(LeftMotor1)
	m.p.Output()

	m.q = rpio.Pin(LeftMotor2)
	m.q.Output()

	m.a = rpio.Pin(RightMotor1)
	m.a.Output()

	m.b = rpio.Pin(RightMotor2)
	m.b.Output()

	return &m
}

// stop both motors
func (m Motors) Stop() {
	m.p.Low()
	m.q.Low()

	m.a.Low()
	m.b.Low()
}

// move forward at speed, 0 <= speed <= 100
func (m Motors) Forward(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	m.p.High()
	m.q.Low()

	m.a.High()
	m.b.Low()
}

// move backwards at speed, 0 <= speed <= 100
func (m Motors) Reverse(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	m.p.Low()
	m.q.High()

	m.a.Low()
	m.b.High()
}

// spin left (sets motors to turn at opposite directions at speed)
// 0 <= speed <= 100
func (m Motors) SpinLeft(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	m.p.High()
	m.q.Low()

	m.a.Low()
	m.b.High()
}

// spin left (sets motors to turn at opposite directions at speed)
// 0 <= speed <= 100
func (m Motors) SpinRight(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	m.p.Low()
	m.q.High()

	m.a.High()
	m.b.Low()
}

// TODO: moves forwards in an arc by setting different speeds
// 0 <= leftSpeed,rightSpeed <= 100
func (m Motors) TurnForward(leftSpeed uint8, rightSpeed uint8) {
	if leftSpeed > 100 || leftSpeed < 0 || rightSpeed > 100 || rightSpeed < 0 {
		fmt.Println("speed out of range")
		return
	}

}

// TODO: moves backwards in an arc by setting different speeds
// 0 <= leftSpeed,rightSpeed <= 100
func (m Motors) TurnReverse(leftSpeed uint8, rightSpeed uint8) {
	if leftSpeed > 100 || leftSpeed < 0 || rightSpeed > 100 || rightSpeed < 0 {
		fmt.Println("speed out of range")
		return
	}

}
