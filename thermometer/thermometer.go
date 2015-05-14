package thermometer

import (
	"io/ioutil"
	"regexp"
)

const (
	baseDir = "C:/sys/bus/w1/devices/"
)

// Thermometer holds thermometer sensor information
type Thermometer struct {
	ID string
}

// New creates a new Thermometer struct
func New() *Thermometer {
	t := &Thermometer{}
	t.discover()

	return t
}

// discover applies the ID of the thermometer sensor in the devices directory
func (t *Thermometer) discover() {
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

func isDeviceName(s string) bool {
	r, err := regexp.Compile(`\d{2}-\w{12}`)

	if err != nil {
		panic(err)
	}

	return r.MatchString(s) == true
}
