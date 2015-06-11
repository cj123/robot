package initio

import (
	"fmt"
	"io/ioutil"
	"os/exec" // sorry :(
)

const (
	Pan  = 0
	Tilt = 1
)

var servosActive = false

func SetServo(servo int, degrees int) {
	if !servosActive {
		StartServos()
	}

	pinServod(servo, degrees)
}

// stop the servos
func StopServos() {
	stopServod()
}

// start the servos
func StartServos() {
	startServod()
}

func startServod() error {
	// run the command
	cmd := exec.Command("./servod", "--pcm", "--idle-timeout=20000", `--p1pins="18,22"`, ">/dev/null")

	err := cmd.Start()

	if err != nil {
		return err
	}

	err = cmd.Wait()

	if err != nil {
		return err
	}

	servosActive = true

	return err
}

func stopServod() error {
	cmd := exec.Command("sudo", "pkill", "-f", "servod")
	err := cmd.Start()

	if err != nil {
		return err
	}

	err = cmd.Wait()

	if err != nil {
		return err
	}

	servosActive = false

	return err
}

func pinServod(pin int, degrees int) {
	pinString := fmt.Sprintf("%d=%d", pin, 50+((90-degrees)*200/180))

	err := ioutil.WriteFile("/dev/servoblaster", []byte(pinString), 0644)

	if err != nil {
		panic(err)
	}
}
