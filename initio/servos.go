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
	fmt.Println("Starting servos...")
	err := startServod()

	if err != nil {
		panic(err)
	}
}

func startServod() error {

	// run the command ./servod --pcm --idle-timeout=20000 --p1pins="18,22"
	cmd := exec.Command("sudo", "sh", "-c", "servod --pcm --idle-timeout=20000 --p1pins=\"18,22\"")
	err := cmd.Run()

	fmt.Println("Starting servod")

	if err != nil {
		fmt.Println(err)
		panic(err)
		return err
	}

	servosActive = true

	return err
}

func stopServod() error {
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

func pinServod(pin int, degrees int) {
	pinString := fmt.Sprintf("%d=%d\n", pin, 50+((90-degrees)*200/180))
	fmt.Println(pinString)
	err := ioutil.WriteFile("/dev/servoblaster", []byte(pinString), 0644)

	if err != nil {
		panic(err)
	}
}
