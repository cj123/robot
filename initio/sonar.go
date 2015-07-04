package initio

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

const (
	// the pin for the ultrasound sensor
	sonarPin = 14

	// speed of sound (cm/s)
	speedOfSound = 34029
)

type Sonar struct{}

func NewSonar() *Sonar {
	return &Sonar{}
}

// return the distance in cm to the nearest reflecting object
// 0 == no object
func (s Sonar) GetDistance() int {
	// setup sonar to be output
	pin := rpio.Pin(sonarPin)
	pin.Output()

	// output a 10us pulse
	pin.High()
	time.Sleep(10 * time.Microsecond)
	pin.Low()

	pin.Input()

	// the start time
	start := time.Now()
	count := time.Now()

	for pin.Read() == rpio.Low && time.Now().Sub(count) < 10*time.Millisecond {
		start = time.Now()
	}

	count = time.Now()
	stop := count

	// wait until the pin read is high
	for pin.Read() == rpio.High && time.Now().Sub(count) < 10*time.Millisecond {
		stop = time.Now()
	}

	// calculate the distance (speed = distance / time, distance = speed * time)
	// ensuring to half the result as this was distance there and back
	dist := stop.Sub(start) * speedOfSound / 2.0 / time.Second

	return int(dist)
}
