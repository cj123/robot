# initio.go

the initio package (as provided [here](http://4tronix.co.uk/initio)), but rewritten in go

**basic implementation, not yet complete.**

## currently working

* ir sensors (right and left, line right and line left)
* ultrasound sensor
* servos (dependency on servod)
* motors (basic implementation)

## todo

* PWM on motors
* arc turns
* speed measurement on the motors

## note

expects `servod` to be in your PATH on the pi :) (see ../setup.sh)