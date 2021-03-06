// +build !linux,!arm

package initio

import (
	"encoding/json"
)

type IR struct{}

func NewIR() *IR {
	return &IR{}
}

// get the state of an ir sensor given by its name
func (ir IR) getIR(sensor string) bool {

	irB, _, err := makeRequest("/api/ir")

	if err != nil {
		panic(err)
	}

	// decode into json
	var output map[string]bool

	err = json.Unmarshal(irB, &output)

	if err != nil {
		panic(err)
	}

	return output[sensor]
}

// returns state of Left IR Obstacle sensor
func (ir IR) Left() bool {
	return ir.getIR("frontleft")
}

// returns state of Right IR Obstacle sensor
func (ir IR) Right() bool {
	return ir.getIR("frontright")
}

// returns state of the Back Right IR Obstacle sensor
//  note: this is an addition i've made myself, not one that is provided
//  as part of the default kit
func (ir IR) BackRight() bool {
	return ir.getIR("backright")
}

// returns state of the Back Left IR Obstacle sensor
//  note: this is an addition i've made myself, not one that is provided
//  as part of the default kit
func (ir IR) BackLeft() bool {
	return ir.getIR("backleft")
}

// returns true if any of the Obstacle sensors are triggered
func (ir IR) All() bool {
	return ir.Right() || ir.Left() || ir.BackRight() || ir.BackLeft()
}

// returns state of Left IR Line sensor
func (ir IR) LeftLine() bool {

	return false // TODO: stubbed
}

// returns state of Right IR Line sensor
func (ir IR) RightLine() bool {

	return false // TODO: stubbed
}
