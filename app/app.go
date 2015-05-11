package app

import (
	"fmt"

	"github.com/danhardman/officr/thermometer"
)

const (
	flux = 0.6
)

var desiredTemp = 20.0

// Start starts the application
func Start() {
	t := thermometer.New()
	t.Discover()

	//TODO: Turn off heating here

	heating := false
	var ct, dt float64

	for true {
		ct = t.Temperature()
		dt = GetDesiredTemperature()
		fmt.Printf("c: %v | d: %v\n", ct, dt)

		if heating {
			if (ct - flux) >= dt {
				DecreaseTemperature()
				heating = false
			}
		} else {
			if (ct + flux) <= dt {
				IncreaseTemperature()
				heating = true
			}
		}
	}
}

// GetDesiredTemperature gets the desired temperature set by the controller
func GetDesiredTemperature() float64 {
	return desiredTemp
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
