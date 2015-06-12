package initio

import (
	"github.com/stianeikeland/go-rpio"
)

type IR struct{}

// the ir sensors
const (
	irFrontLeft  = 4  // 7 but using the mapping from go-rpio
	irFrontRight = 17 // 11
	irLineLeft   = 18 // 12
	irLineRight  = 21 // 13
)

// makes things a bit nicer
func NewIR() *IR {
	return &IR{}
}

// returns state of Left IR Obstacle sensor
func (ir IR) Left() bool {
	pin := rpio.Pin(irFrontLeft)
	pin.Input()

	return pin.Read() == rpio.Low
}

// returns state of Right IR Obstacle sensor
func (ir IR) Right() bool {
	pin := rpio.Pin(irFrontRight)
	pin.Input()

	return pin.Read() == rpio.Low
}

// returns true if any of the Obstacle sensors are triggered
func (ir IR) All() bool {
	return ir.Right() || ir.Left()
}

// returns state of Left IR Line sensor
func (ir IR) LeftLine() bool {
	pin := rpio.Pin(irLineLeft)
	pin.Input()

	return pin.Read() == rpio.Low
}

// returns state of Right IR Line sensor
func (ir IR) RightLine() bool {
	pin := rpio.Pin(irLineRight)
	pin.Input()

	return pin.Read() == rpio.Low
}
