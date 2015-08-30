package initio

import (
	"github.com/stianeikeland/go-rpio"
)

type IR struct{}

// makes things a bit nicer
func NewIR() *IR {
	return &IR{}
}

// get if the given pin is low (triggered, since active low)
func (ir IR) pinRead(reqPin int) bool {
	pin := rpio.Pin(reqPin)
	pin.Input()

	return pin.Read() == rpio.Low
}

// returns state of Left IR Obstacle sensor
func (ir IR) Left() bool {
	return ir.pinRead(irFrontLeft)
}

// returns state of Right IR Obstacle sensor
func (ir IR) Right() bool {
	return ir.pinRead(irFrontRight)
}

// returns state of the Back Right IR Obstacle sensor
//  note: this is an addition i've made myself, not one that is provided
//  as part of the default kit
func (ir IR) BackRight() bool {
	return ir.pinRead(irBackRight)
}

// returns state of the Back Left IR Obstacle sensor
//  note: this is an addition i've made myself, not one that is provided
//  as part of the default kit
func (ir IR) BackLeft() bool {
	return ir.pinRead(irBackLeft)
}

// returns true if any of the Obstacle sensors are triggered
func (ir IR) All() bool {
	return ir.Right() || ir.Left() || ir.BackRight() || ir.BackLeft()
}

// returns state of Left IR Line sensor
func (ir IR) LeftLine() bool {
	return ir.pinRead(irLineLeft)
}

// returns state of Right IR Line sensor
func (ir IR) RightLine() bool {
	return ir.pinRead(irLineRight)
}
