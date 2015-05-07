package app

import (
	"fmt"
	"time"
)

const (
	baseDir = "/sys/devices/"
	flux    = 0.6
)

type Thermometer struct {
	file string
}

var heating = false
var currentTemp = 15.0
var desiredTemp = 20.0

// Start starts the application
func Start() {
	for true {
		ct := GetCurrentTemperature()
		dt := GetDesiredTemperature()

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
		time.Sleep(100 * time.Millisecond)
	}
}

func NewThermometer() *Thermometer {
	return &Thermometer{}
}

// GetCurrentTemperature gets the current temperature from the temperature sensor
func GetCurrentTemperature() float64 {
	if heating {
		currentTemp = currentTemp + 0.1
	} else {
		currentTemp = currentTemp - 0.1
	}
	fmt.Println(currentTemp)
	return currentTemp
}

// GetDesiredTemperature gets the desired temperature set by the controller
func GetDesiredTemperature() float64 {
	fmt.Println(desiredTemp)
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
