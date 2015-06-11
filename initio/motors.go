package initio

// stop both motors
func Stop() {

}

// move forward at speed, 0 <= speed <= 100
func Forward(speed uint8) {

}

// move backwards at speed, 0 <= speed <= 100
func Reverse(speed uint8) {

}

// spin left (sets motors to turn at opposite directions at speed)
// 0 <= speed <= 100
func SpinLeft(speed uint8) {

}

// spin left (sets motors to turn at opposite directions at speed)
// 0 <= speed <= 100
func SpinRight(speed uint8) {

}

//  moves forwards in an arc by setting different speeds
// 0 <= leftSpeed,rightSpeed <= 100
func TurnForward(leftSpeed uint8, rightSpeed uint8) {

}

// moves backwards in an arc by setting different speeds
// 0 <= leftSpeed,rightSpeed <= 100
func TurnReverse(leftSpeed uint8, rightSpeed uint8) {

}
