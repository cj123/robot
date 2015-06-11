package initio

import (
	"testing"
	"time"
)

const interval = time.Second

func TestForwards(t *testing.T) {
	t.Log("Testing forwards")

	for i := 0; i < 3; i++ {
		Forward(0)
		time.Sleep(interval)

		Stop()
		time.Sleep(interval)
	}

	Stop()
}

func TestReverse(t *testing.T) {
	t.Log("Testing reverse")

	for i := 0; i < 3; i++ {
		Reverse(0)
		time.Sleep(interval)

		Stop()
		time.Sleep(interval)
	}

	Stop()
}

func TestRight(t *testing.T) {
	t.Log("Testing right")

	for i := 0; i < 3; i++ {
		SpinRight(0)
		time.Sleep(interval)

		Stop()
		time.Sleep(interval)
	}

	Stop()
}

func TestLeft(t *testing.T) {
	t.Log("Testing left")

	for i := 0; i < 3; i++ {
		SpinLeft(0)
		time.Sleep(interval)

		Stop()
		time.Sleep(interval)
	}

	Stop()
	Cleanup()
}
