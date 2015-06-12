package initio

import (
	"testing"
	"time"
)

func TestServos(t *testing.T) {
	// move them around a bit
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
