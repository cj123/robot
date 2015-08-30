package initio

import (
	"fmt"
)

type Motors struct {
	// the motor points
	a, b, p, q *PWMPin
}

func NewMotor() *Motors {
	m := Motors{}

	// initialise the pins
	m.p = NewPWMPin(LeftMotor1)
	m.p.Output()

	m.q = NewPWMPin(LeftMotor2)
	m.q.Output()

	m.a = NewPWMPin(RightMotor1)
	m.a.Output()

	m.b = NewPWMPin(RightMotor2)
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

	m.p.pwm(speed)
	m.q.Low()

	m.a.pwm(speed)
	m.b.Low()
}

// move backwards at speed, 0 <= speed <= 100
func (m Motors) Reverse(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	m.p.Low()
	m.q.pwm(speed)

	m.a.Low()
	m.b.pwm(speed)
}

// spin left (sets motors to turn at opposite directions at speed)
// 0 <= speed <= 100
func (m Motors) SpinLeft(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	m.p.pwm(speed)
	m.q.Low()

	m.a.Low()
	m.b.pwm(speed)
}

// spin left (sets motors to turn at opposite directions at speed)
// 0 <= speed <= 100
func (m Motors) SpinRight(speed uint8) {
	if speed > 100 || speed < 0 {
		fmt.Println("speed out of range")
		return
	}

	m.p.Low()
	m.q.pwm(speed)

	m.a.pwm(speed)
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
