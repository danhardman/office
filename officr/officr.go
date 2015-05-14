package main

import (
	"github.com/danhardman/officr/server"
	"github.com/danhardman/officr/thermometer"
	"github.com/danhardman/officr/thermostat"
)

func main() {
	t := thermometer.New()
	r := make(chan *thermostat.Reading)

	go thermostat.Reader(t.ID, r)
	go thermostat.Regulate(r)

	go server.Router()

	select {}
}
