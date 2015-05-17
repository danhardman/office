package thermostat

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

const (
	flux = 0.6
)

var desiredTemp = 20.0

// Regulate monitors the current temperature and regulates the heating
func Regulate(r <-chan *Reading) {
	// TODO: Turn off heating here

	heating := false
	var dt float64

	for cr := range r {
		dt = GetDesiredTemperature()
		fmt.Printf("c: %v | d: %v\n", cr.Temperature, dt)

		if heating {
			if (cr.Temperature - flux) >= dt {
				DecreaseTemperature()
				heating = false
			}
		} else {
			if (cr.Temperature + flux) <= dt {
				IncreaseTemperature()
				heating = true
			}
		}
	}
}

// GetDesiredTemperature gets the desired temperature set by the controller
func GetDesiredTemperature() float64 {
	f, err := filepath.Abs("../temperature")

	if err != nil {
		panic(err)
	}

	dt, err := ioutil.ReadFile(f)

	if err != nil {
		panic(err)
	}

	t, err := strconv.ParseFloat(string(dt[:]), 64)

	if err != nil {
		panic(err)
	}

	return t
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
