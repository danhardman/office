package thermostat

import (
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	baseDir = "C:/sys/bus/w1/devices/"
)

type Reading struct {
	TID         string
	Temperature float64
}

// Reader reads the provided thermometer to get its temperature
func Reader(ID string, c chan<- *Reading) {
	for {
		r := &Reading{
			TID: ID,
		}

		o, err := ioutil.ReadFile(baseDir + ID + "/w1_slave")

		if err != nil {
			panic(err)
		}

		ts := strings.Split(string(o), "\n")

		if !strings.Contains(ts[0], "YES") {
			break
		} else {
			ti, err := strconv.Atoi(ts[1][len(ts[1])-5 : len(ts[1])])

			if err != nil {
				panic(err)
			}

			r.Temperature = float64(ti / 1000)
			c <- r
		}
	}
}
