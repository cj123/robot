package initio

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
)

const (
	// right motor pins
	RightMotor1 = 8 // 24
	RightMotor2 = 7 // 26

	// left motor pins
	LeftMotor1 = 10 // 19
	LeftMotor2 = 9  // 21
)

// the motor points
var a, b, p, q rpio.Pin

func init() {
	// initialise the pins
	p = rpio.Pin(LeftMotor1)
	p.Output()

	q = rpio.Pin(LeftMotor2)
	q.Output()

	a = rpio.Pin(RightMotor1)
	a.Output()

	b = rpio.Pin(RightMotor2)
	b.Output()
}

// stop both motors
func Stop() {
	p.Low()
	q.Low()

	a.Low()
	b.Low()
}

// move forward at speed, 0 <= speed <= 100
func Forward(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	p.High()
	q.Low()

	a.High()
	b.Low()
}

// move backwards at speed, 0 <= speed <= 100
func Reverse(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	p.Low()
	q.High()

	a.Low()
	b.High()
}

// spin left (sets motors to turn at opposite directions at speed)
// 0 <= speed <= 100
func SpinLeft(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	p.High()
	q.Low()

	a.Low()
	b.High()
}

// spin left (sets motors to turn at opposite directions at speed)
// 0 <= speed <= 100
func SpinRight(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	p.Low()
	q.High()

	a.High()
	b.Low()
}

// TODO:
//  moves forwards in an arc by setting different speeds
// 0 <= leftSpeed,rightSpeed <= 100
func TurnForward(leftSpeed uint8, rightSpeed uint8) {
	if leftSpeed > 100 || leftSpeed < 0 || rightSpeed > 100 || rightSpeed < 0 {
		fmt.Println("speed out of range")
		return
	}

}

// TODO:
// moves backwards in an arc by setting different speeds
// 0 <= leftSpeed,rightSpeed <= 100
func TurnReverse(leftSpeed uint8, rightSpeed uint8) {
	if leftSpeed > 100 || leftSpeed < 0 || rightSpeed > 100 || rightSpeed < 0 {
		fmt.Println("speed out of range")
		return
	}

}
