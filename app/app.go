package app

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	BaseDir = "/device/directory/"
)

type Thermometer struct {
	device      string
	temperature float64
}

func Start() {

}

func GetTemperature(d string) *Thermometer {
	data, err := ioutil.ReadFile(BaseDir + d + "/w1_slave")

	if err != nil {
		panic(err)
	}

	ts := strings.Split(string(data), "\n")
	i, err := strconv.Atoi(ts[1][len(ts[1])-5 : len(ts[1])])

	therm := &Thermometer{
		device:      d,
		temperature: float64(i) / 1000,
	}

	return therm
}
