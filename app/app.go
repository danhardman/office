package app

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	wiggle = 1.5
)

// Start starts the application
func Start() {
	for true {
		ct := GetCurrentTemperature()
		dt := GetDesiredTemperature()

		if math.Abs(ct-dt) > wiggle {
			if ct > dt {
				DecreaseTemperature()
			} else if ct < dt {
				IncreaseTemperature()
			}
		}
		time.Sleep(10 * time.Second)
	}
}

// GetCurrentTemperature gets the current temperature from the temperature sensor
func GetCurrentTemperature() float64 {
	ct := rand.Float64() * 10
	fmt.Println(ct)
	return ct
}

// GetDesiredTemperature gets the desired temperature set by the controller
func GetDesiredTemperature() float64 {
	dt := rand.Float64() * 10
	fmt.Println(dt)
	return dt
}

// DecreaseTemperature decreases the temperature by decreasing heating and
// increasing cooling
func DecreaseTemperature() {
	fmt.Println("Decreasing temp")
}

// IncreaseTemperature increases the temperature by increasing the heating and
// decreasing cooling
func IncreaseTemperature() {
	fmt.Println("Increasing temp")
}
