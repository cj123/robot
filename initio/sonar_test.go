package initio

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDistance(t *testing.T) {
	s := NewSonar()

	for i := 0; i < 10; i++ {
		fmt.Printf("Distance: %dcm\n", s.GetDistance())
		time.Sleep(2 * time.Second)
	}

	// clean up at the end
	Cleanup()
}
