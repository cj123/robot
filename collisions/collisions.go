package collisions

import (
	"fmt"
	"github.com/cj123/robot/initio"
	"math"
	"time"
)

func init() {
	inMotion = false

	// make all reading values -1 to begin with
	for i := 0; i < len(readings); i++ {
		averageReadings[i] = -1
		for j := 0; j < len(readings[i]); j++ {
			readings[i][j] = -1
		}
	}
}

// the readings from the sensors at given tilt/pan angles
// as there are +90 -> -90 degrees on each centre, that gives 180*180
// possible readings. How quick these will happen, we don't know let's see
var readings [180][180]int

// the average readings of each sensor pan location (averaged tilts)
var averageReadings [180]int

// because we're using a 2D array, we need to adjust the index to be in the range
// 0 < 180 rather than -90 < 90.
func getIndex(i int) int {
	return i + 90
}

const (
	PAN_STEP  = 10
	TILT_STEP = 20
)

func MakeReadings() {

	start := time.Now()
	numReadings := 0

	for i := -90; i < 90; i += PAN_STEP {
		// move the servo pan position
		initio.SetServo(initio.Pan, i)

		for j := -60; j < 60; j += TILT_STEP {
			dist := initio.GetDistance()

			//fmt.Printf("(%d, %d) = %d\n", i, j, dist)

			// store reading
			readings[getIndex(i)][getIndex(j)] = dist

			// move the servo tilt position
			initio.SetServo(initio.Tilt, j)

			time.Sleep(50 * time.Millisecond)
			numReadings++
		}
	}

	end := time.Now()

	fmt.Printf("Total time taken: %fs\n", end.Sub(start).Seconds())
	fmt.Printf("Number of readings: %d\n", numReadings)
}

// get the position for the maximum distance, as well as that distance
// note that maxKey is returned as servo position, not array index
// (use getIndex() to get the array index)
func GetMaximumDistance() (maxKey int, maxVal int) {

	// if not in motion
	if !inMotion {
		// do quick readings
		MakeReadings()
	} else {
		// do a full scan
		MakeReadings()
	}

	for i := 0; i < len(readings); i++ {
		sum := 0
		numTaken := 0
		for j := 0; j < len(readings[i]); j++ {
			if readings[i][j] == -1 {
				continue
			}

			// average all the y values into an array
			sum += readings[i][j]

			// for the denominator
			numTaken++

		}

		// avoid division by zero
		if numTaken != 0 {
			averageReadings[i] = sum / numTaken
		}
	}

	// find the maximum value
	maxVal = averageReadings[0]
	maxKey = 0

	for index, val := range averageReadings {
		//fmt.Printf("%d = %f\n", index, val)
		if val != -1 && val != 0 && val > maxVal {

			maxVal = val
			maxKey = index
		}
	}

	return maxKey - 90, maxVal
}

// are we in motion?
var inMotion bool

const (
	DIR_LEFT = iota
	DIR_FRONT
	DIR_RIGHT
	DIR_UNKNOWN
)

var directionNames = [...]string{"left", "front", "right", "unknown :("}

// given a servo point, get the general direction we should turn
func GetDirection(servoPos int) int {
	if servoPos < 35 && servoPos > 5 { // +/- 15 on the FRONT position (20)
		return DIR_FRONT
	} else if servoPos >= 35 {
		return DIR_RIGHT
	} else if servoPos <= 5 {
		return DIR_LEFT
	} else {
		return DIR_UNKNOWN
	}
}

func GetDirectionName(dir int) string {
	return directionNames[dir]
}

// given a turn degree, ESTIMATE (for now) how long that will take to turn
func getTimeForTurn(degrees int) time.Duration {
	// ok, so it's nowhere near perfect, but i think it takes about
	// 1 second for the robot to do a 90 degree turn
	// so let's try this
	return time.Duration(math.Abs(float64(initio.DEFAULT_VAL-degrees))) * (time.Second / 90)
}

const ROBOT_SPEED = 4 // cm/s

// speed = distance / time
// speed * time = distance
// time = distance / speed
func getTimeToMoveForwards(distance int) time.Duration {
	// time = distance * speed
	return time.Duration(distance/ROBOT_SPEED) * time.Second
}

func DoFunkyCollisionAvoidance() bool {

	ir := initio.IR{}

	for {

		// check ir
		if ir.Left() && ir.Right() {
			// reverse
			initio.Reverse(0)
			time.Sleep(20 * time.Millisecond)
			initio.Stop()
			continue
		}

		maxKey, maxVal := GetMaximumDistance()

		direction := GetDirection(maxKey)

		fmt.Printf("Maximum value: %d at key %d\n", maxVal, maxKey)
		fmt.Printf("So we should probably turn %s\n", GetDirectionName(direction))

		fmt.Println("Moving servo to that position...")

		initio.SetServo(initio.Pan, maxKey)
		initio.SetServo(initio.Tilt, initio.DEFAULT_VAL)

		// maybe use a switch here
		// find the turn degrees required to do it
		if direction == DIR_FRONT {
			// go forwards
			initio.Forward(0) // speed still needs implementing
			time.Sleep(getTimeToMoveForwards(maxVal))
		} else if direction == DIR_LEFT {
			initio.SpinLeft(0)
			t := getTimeForTurn(maxKey)
			fmt.Println("I should turn for", t)
			time.Sleep(t)
		} else if direction == DIR_RIGHT {
			initio.SpinRight(0)
			t := getTimeForTurn(maxKey)
			fmt.Println("I should turn for", t)
			time.Sleep(t)
		} else {
			// uh we done fuckup
			fmt.Println("broken")
		}

		initio.Stop()

	}

	return true
}
