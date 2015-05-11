package thermometer

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	baseDir = "C:/officr/sys/bus/w1/devices/"
)

// Thermometer holds thermometer sensor information
type Thermometer struct {
	ID string
}

// New creates a new Thermometer struct
func New() *Thermometer {
	return &Thermometer{}
}

// Discover applies the ID of the thermometer sensor in the devices directory
func (t *Thermometer) Discover() {
	devices, err := ioutil.ReadDir(baseDir)

	if err != nil {
		panic(err)
	}

	match := false
	for _, d := range devices {
		if isDeviceName(d.Name()) {
			t.ID = d.Name()
			match = true
			break
		}
	}

	if !match {
		panic("No thermometer device found!")
	}
}

// Temperature returns the temperature read from the Thermometer sensor
func (t *Thermometer) Temperature() float64 {
	o, err := ioutil.ReadFile(baseDir + t.ID + "/w1_slave")

	if err != nil {
		panic(err)
	}

	ts := strings.Split(string(o), "\n")
	ti, err := strconv.Atoi(ts[1][len(ts[1])-5 : len(ts[1])])

	if err != nil {
		panic(err)
	}

	return float64(ti / 1000)
}

func isDeviceName(s string) bool {
	r, err := regexp.Compile(`\d{2}-\w{12}`)

	if err != nil {
		panic(err)
	}

	return r.MatchString(s) == true
}
