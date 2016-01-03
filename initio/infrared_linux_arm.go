package initio

import (
	"github.com/stianeikeland/go-rpio"
)

// IR sensors
type IR struct{}

// NewIR returns an instance of IR
func NewIR() *IR {
	return &IR{}
}

// get if the given pin is low (triggered, since active low)
func (ir IR) pinRead(reqPin int) bool {
	pin := rpio.Pin(reqPin)
	pin.Input()

	return pin.Read() == rpio.Low
}

// Left returns state of Left IR Obstacle sensor
func (ir IR) Left() bool {
	return ir.pinRead(irFrontLeft)
}

// Right returns state of Right IR Obstacle sensor
func (ir IR) Right() bool {
	return ir.pinRead(irFrontRight)
}

// BackRight returns state of the Back Right IR Obstacle sensor
// note: this is an addition i've made myself, not one that is provided
// as part of the default kit
func (ir IR) BackRight() bool {
	return ir.pinRead(irBackRight)
}

// BackLeft returns state of the Back Left IR Obstacle sensor
// note: this is an addition i've made myself, not one that is provided
// as part of the default kit
func (ir IR) BackLeft() bool {
	return ir.pinRead(irBackLeft)
}

// All returns true if any of the Obstacle sensors are triggered
func (ir IR) All() bool {
	return ir.Right() || ir.Left() || ir.BackRight() || ir.BackLeft()
}

// LeftLine returns state of Left IR Line sensor
func (ir IR) LeftLine() bool {
	return ir.pinRead(irLineLeft)
}

// RightLine returns state of Right IR Line sensor
func (ir IR) RightLine() bool {
	return ir.pinRead(irLineRight)
}
