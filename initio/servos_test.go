package initio

import (
	"testing"
	"time"
)

func TestServos(t *testing.T) {
	SetServo(Pan, 10)
	time.Sleep(5 * time.Second)
	SetServo(Pan, 20)
	time.Sleep(5 * time.Second)
	SetServo(Pan, -20)

	time.Sleep(5 * time.Second)
	SetServo(Tilt, 10)
	time.Sleep(5 * time.Second)
	SetServo(Tilt, 20)
	time.Sleep(5 * time.Second)
	SetServo(Tilt, -20)

	Cleanup()
}
