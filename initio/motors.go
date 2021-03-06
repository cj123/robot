// +build !linux,!arm

package initio

import (
	"fmt"
)

type Motors struct{}

func NewMotor() *Motors {
	return &Motors{}
}

// stop both motors
func (m Motors) Stop() {
	makeRequest("/api/motors/stop")
}

// move forward at speed, 0 <= speed <= 100
func (m Motors) Forward(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	makeRequest(fmt.Sprintf("/api/motors/forwards/%d", speed))
}

// move backwards at speed, 0 <= speed <= 100
func (m Motors) Reverse(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	makeRequest(fmt.Sprintf("/api/motors/reverse/%d", speed))
}

// spin left (sets motors to turn at opposite directions at speed)
// 0 <= speed <= 100
func (m Motors) SpinLeft(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	makeRequest(fmt.Sprintf("/api/motors/left/%d", speed))
}

// spin left (sets motors to turn at opposite directions at speed)
// 0 <= speed <= 100
func (m Motors) SpinRight(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	makeRequest(fmt.Sprintf("/api/motors/right/%d", speed))
}

// TODO: moves forwards in an arc by setting different speeds
// 0 <= leftSpeed,rightSpeed <= 100
func (m Motors) TurnForward(leftSpeed uint8, rightSpeed uint8) {
	if leftSpeed > 100 || leftSpeed < 0 || rightSpeed > 100 || rightSpeed < 0 {
		fmt.Println("speed out of range")
		return
	}

	// TODO: stubbed
}

// TODO: moves backwards in an arc by setting different speeds
// 0 <= leftSpeed,rightSpeed <= 100
func (m Motors) TurnReverse(leftSpeed uint8, rightSpeed uint8) {
	if leftSpeed > 100 || leftSpeed < 0 || rightSpeed > 100 || rightSpeed < 0 {
		fmt.Println("speed out of range")
		return
	}

	// TODO: stubbed
}
