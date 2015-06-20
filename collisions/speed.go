package collisions

import (
	"github.com/cj123/robot/initio"
	"log"
	"math"
	"time"
)

// a complete estimation for now
const ROBOT_SPEED = 30 // cms^-1

// index corresponds to direction name
var directionNames = [...]string{"left", "front", "right", "unknown :("}

// given a servo point, get the general direction we should turn
func getDirection(servoPos int) int {
	if servoPos < 40 && servoPos > 0 { // +/- 20 on the FRONT position (20)
		return DIR_FRONT
	} else if servoPos >= 40 {
		return DIR_RIGHT
	} else if servoPos <= 0 {
		return DIR_LEFT
	} else {
		return DIR_UNKNOWN
	}
}

func getDirectionName(dir int) string {
	return directionNames[dir]
}

// given a turn degree, ESTIMATE (for now) how long that will take to turn
func getTimeForTurn(degrees int) time.Duration {
	// ok, so it's nowhere near perfect, but i think it takes about
	// 1 second for the robot to do a 90 degree turn
	// so let's try this
	return time.Duration(math.Abs(float64(initio.DEFAULT_VAL-degrees))) * (time.Second / 90)
}

// speed = distance / time
// speed * time = distance
// time = distance / speed
func getTimeToMoveForwards(distance int) time.Duration {
	// time = distance * speed
	t := time.Duration(distance/ROBOT_SPEED) * time.Second

	log.Println("My speed is", ROBOT_SPEED)
	log.Println("The distance I read is", distance)
	log.Println("I think i should move forward for", t)

	return t
}
