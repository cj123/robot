package initio

import (
	"testing"
	"time"
)

const interval = time.Second
const testSpeed = 100

var m *Motors

func init() {
	m = NewMotor()
}

func TestForwards(t *testing.T) {
	t.Log("Testing forwards")

	for i := 0; i < 3; i++ {
		m.Forward(testSpeed)
		time.Sleep(interval)

		m.Stop()
		time.Sleep(interval)
	}

	m.Stop()
}

func TestReverse(t *testing.T) {
	t.Log("Testing reverse")

	for i := 0; i < 3; i++ {
		m.Reverse(testSpeed)
		time.Sleep(interval)

		m.Stop()
		time.Sleep(interval)
	}

	m.Stop()
}

func TestRight(t *testing.T) {
	t.Log("Testing right")

	for i := 0; i < 3; i++ {
		m.SpinRight(testSpeed)
		time.Sleep(interval)

		m.Stop()
		time.Sleep(interval)
	}

	m.Stop()
}

func TestLeft(t *testing.T) {
	t.Log("Testing left")

	for i := 0; i < 3; i++ {
		m.SpinLeft(testSpeed)
		time.Sleep(interval)

		m.Stop()
		time.Sleep(interval)
	}

	m.Stop()
	Cleanup()
}
