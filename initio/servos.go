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

// set a servo to a certain angle
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
		panic(err)
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
		panic(err)
		return err
	}

	servosActive = false

	return err
}

// apply servo change to the pin
func pinServod(pin int, degrees int) {
	pinString := fmt.Sprintf("%d=%d\n", pin, 50+((90-degrees)*200/180))
	fmt.Println(pinString)
	err := ioutil.WriteFile("/dev/servoblaster", []byte(pinString), 0644)

	if err != nil {
		panic(err)
	}
}
