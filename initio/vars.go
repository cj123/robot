package initio

const (
	// servos
	Pan         = 0
	Tilt        = 1
	DEFAULT_VAL = 0

	MAX_VAL = 90
	MIN_VAL = -90

	// motors
	RightMotor1 = 8 // right motor pins 24
	RightMotor2 = 7 // 26

	LeftMotor1 = 10 // left motor pins: 19
	LeftMotor2 = 9  // 21

	// infrared sensors
	irFrontLeft  = 4  // 7 but using the mapping from go-rpio
	irFrontRight = 17 // 11
	irBackLeft   = 23 // !! TODO
	irBackRight  = 22 // !! TODO
	irLineLeft   = 18 // 12
	irLineRight  = 21 // 13
)

var servoNames = [...]string{"pan", "tilt"}
