// +build !linux,!arm

package initio

import (
	"fmt"
	"strconv"
	"strings"
)

type Sonar struct{}

func NewSonar() *Sonar {
	return &Sonar{}
}

// return the distance in cm to the nearest reflecting object
// 0 == no object
func (s Sonar) GetDistance() int {
	str, _, err := makeRequest("/api/sonar/distance")

	if err != nil {
		panic(err)
	}

	num := strings.Replace(string(str), "\n", "", -1)

	dist, err := strconv.ParseInt(num, 10, 10)

	if err != nil {
		panic(err)
	}

	return int(dist)
}
