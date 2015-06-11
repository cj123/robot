package initio

import (
	"os/exec" // sorry :(
)

var servosActive = false

func SetServo(servo int, degrees int) {
	if !servosActive {
		StartServos()
	}

	pinServod(servo, degrees)
}

func StopServos() {
	stopServod()
}

func StartServos() {
	startServod()
}

func startServod() {
	// run the command
	cmd := exec.Command("./servod", "--pcm", "--idle-timeout=20000", `--p1pins="18,22"`, ">/dev/null")

	err := cmd.Start()

	if err != nil {
		panic(err) // for now
	}

	err = cmd.Wait()

	if err != nil {
		panic(err)
	}
	servosActive = true
}

func stopServod() {
	cmd := exec.Command("sudo", "pkill", "-f", "servod")
	err := cmd.Start()

	if err != nil {
		panic(err)
	}

	err = cmd.Wait()

	if err != nil {
		panic(err)
	}

	servosActive = false
}

func pinServod(pin int, degrees int) {

}
