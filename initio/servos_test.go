package initio

import (
	"testing"
	"time"
)

func TestServos(t *testing.T) {

	panServo := NewServo(Pan)
	tiltServo := NewServo(tilt)

	// move them around a bit
	panServo.Set(10)
	time.Sleep(5 * time.Second)
	panServo.Set(20)
	time.Sleep(5 * time.Second)
	panServo.Set(-20)

	time.Sleep(5 * time.Second)
	tiltServo.Set(10)
	time.Sleep(5 * time.Second)
	tiltServo.Set(20)
	time.Sleep(5 * time.Second)
	tiltServo.Set(-20)

	Cleanup()
}
