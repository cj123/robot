package collisions

import (
	"github.com/cj123/robot/initio"
	"log"
	"time"
)

type Reading int

const (
	// the directions
	DIR_LEFT = iota
	DIR_FRONT
	DIR_RIGHT
	DIR_UNKNOWN

	// type of reading
	READING_QUICK Reading = iota
	READING_FULL  Reading = iota

	// the increments for the servos
	PAN_STEP  = 10
	TILT_STEP = 20

	// an adjustment to take from the time calculated on a corner
	CORNER_ADJUST = 25 * time.Millisecond
)

var (
	// the readings from the sensors at given tilt/pan angles
	// as there are +90 -> -90 degrees on each centre, that gives 180*180
	// possible readings. How quick these will happen, we don't know let's see
	readings [180][180]int

	// the average readings of each sensor pan location (averaged tilts)
	averageReadings [180]int

	// are we in motion?
	inMotion bool

	takingMeasurements, foundCollision bool
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

// because we're using a 2D array, we need to adjust the index to be in the range
// 0 < 180 rather than -90 < 90.
func getIndex(i int) int {
	return i + 90
}

// make readings, populating the readings
func MakeReadings(readingType Reading) {

	takingMeasurements = true

	// get the step values
	panStep, tiltStep := PAN_STEP, TILT_STEP

	// if it's a quick read, increase them
	if readingType == READING_QUICK {
		panStep *= 2
		tiltStep *= 3
	}

	start := time.Now()
	numReadings := 0

	for i := -90; i < 90; i += panStep {
		// move the servo pan position
		initio.SetServo(initio.Pan, i)

		for j := -60; j < 90; j += tiltStep {
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

	log.Printf("Total time taken: %fs\n", end.Sub(start).Seconds())
	log.Printf("Number of readings: %d\n", numReadings)

	takingMeasurements = false
}

// get the position for the maximum distance, as well as that distance
// note that maxKey is returned as servo position, not array index
// (use getIndex() to get the array index)
func GetMaximumDistance() (maxKey int, maxVal int) {

	// if not in motion
	if !inMotion {
		// do quick readings
		MakeReadings(READING_QUICK)
	} else {
		// do a full scan
		MakeReadings(READING_FULL)
	}

	// go through all the readings
	for i := 0; i < len(readings); i++ {
		sum := 0
		numTaken := 0

		for j := 0; j < len(readings[i]); j++ {
			// if it's set to not read here,
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

// constantly recheck the IR sensors,
func checkIR() {

	ir := initio.IR{}

	for {
		time.Sleep(10 * time.Microsecond)

		if takingMeasurements {
			log.Println("Taking measurements...")
			continue
		}

		log.Println("Checking for collisions")

		// check ir
		if ir.Left() || ir.Right() {

			log.Println("Found immediate collision. Stopping...")

			foundCollision = true

			initio.Stop() // stop immediately

			// reverse until the objects are gone
			initio.Reverse(0)

			for ir.Left() || ir.Right() {
				time.Sleep(10 * time.Microsecond)

				// block any other actions until we're out of danger
				initio.Reverse(0)
			}

			initio.Stop()

			log.Println("Object moved, continuing...")

			foundCollision = false
		}
	}
}

// start collision avoidance
func Start(run *bool) bool {

	if *run {
		log.Println("Starting collision avoidance...")
	} else {
		log.Println("Collision avoidance initialised, but disabled for now")
	}

	go func() {
		foundCollision = false

		// start the IR sensor checking
		checkIR()
	}()

	for {
		// if we found a collision
		if foundCollision || !*run {
			log.Println("Collision found, no readings to be taken")
			time.Sleep(10 * time.Millisecond)
			continue
		}

		maxKey, maxVal := GetMaximumDistance()
		direction := getDirection(maxKey)

		log.Printf("Maximum value: %d at key %d\n", maxVal, maxKey)
		log.Printf("So we should probably turn %s\n", getDirectionName(direction))

		log.Println("Moving servo to that position...")

		// set the servo to the turn value
		initio.SetServo(initio.Pan, maxKey)

		// set the servo height, because... OCD
		initio.SetServo(initio.Tilt, initio.DEFAULT_VAL)

		// find the turn degrees required to do it
		switch direction {
		case DIR_FRONT:
			// go forwards
			initio.Forward(0) // speed still needs implementing
			inMotion = true
			time.Sleep(getTimeToMoveForwards(maxVal))
			break
		case DIR_LEFT:
			initio.SpinLeft(0)
			t := getTimeForTurn(maxKey)
			log.Println("I should turn for", t)
			inMotion = true
			time.Sleep(t)

			initio.Forward(0) // then keep going

			// get the time to move forward, but take a small amount off
			// because of the corner we just turned
			time.Sleep(getTimeToMoveForwards(maxVal) - CORNER_ADJUST)
			break
		case DIR_RIGHT:
			initio.SpinRight(0)
			t := getTimeForTurn(maxKey)
			log.Println("I should turn for", t)
			inMotion = true
			time.Sleep(t)

			initio.Forward(0) // then keep going

			// get the time to move forward, but take a small amount off
			// because of the corner we just turned
			time.Sleep(getTimeToMoveForwards(maxVal) - CORNER_ADJUST)
			break
		default:
			log.Println("broken")
		}

		// stop moving
		initio.Stop()

		inMotion = false
	}

	return true
}
