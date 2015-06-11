package initio

import (
	"fmt"
	"testing"
	"time"
)

// test all the ir sensors
func TestIR(t *testing.T) {
	ir := NewIR()

	left, right, lineLeft, lineRight := ir.Left(), ir.Right(), ir.LeftLine(), ir.RightLine()

	fmt.Println("| Left  | Right | LLeft | LRight |")
	fmt.Println("----------------------------------")

	fmt.Printf("| %-5t | %-5t | %-5t | %-5t  |", left, right, lineLeft, lineRight)

	for {
		// calculate new values
		newL, newR, newLineL, newLineR := ir.Left(), ir.Right(), ir.LeftLine(), ir.RightLine()

		// if any values change
		if newL != left || newR != right || newLineL != lineLeft || newLineR != lineRight {
			// update the old ones
			left, right, lineLeft, lineRight := newL, newR, newLineL, newLineR

			// print them
			fmt.Printf("| %-5t | %-5t | %-5t | %-5t  |\n", left, right, lineLeft, lineRight)
		}

		time.Sleep(100 * time.Millisecond)
	}

	Cleanup()
}
