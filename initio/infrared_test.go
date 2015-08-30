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
	bLeft, bRight := ir.BackLeft(), ir.BackRight()

	fmt.Println("| Left  | Right | BLeft | BRight | LLeft | LRight |")
	fmt.Println("---------------------------------------------------")

	fmt.Printf("| %-5t | %-5t | %-5t | %-5t | %-5t | %-5t  |\n", left, right, bLeft, bRight, lineLeft, lineRight)

	for i := 0; i < 20; i++ {
		// calculate new values
		newL, newR, newLineL, newLineR := ir.Left(), ir.Right(), ir.LeftLine(), ir.RightLine()
		newbLeft, newbRight := ir.BackLeft(), ir.BackRight()

		// if any values change
		if newL != left || newR != right || newLineL != lineLeft || newLineR != lineRight ||
			newbLeft != bLeft || newbRight != bRight {
			// update the old ones
			left, right, lineLeft, lineRight = newL, newR, newLineL, newLineR
			bLeft, bRight = newbLeft, newbRight
			// print them
			fmt.Printf("| %-5t | %-5t | %-5t | %-5t | %-5t | %-5t  |\n", left, right, bLeft, bRight, lineLeft, lineRight)
		}

		time.Sleep(100 * time.Millisecond)
	}

	Cleanup()
}
