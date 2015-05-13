package main

import (
	"net/http"

	"github.com/danhardman/officr/server"
	"github.com/danhardman/officr/thermostat"
)

func main() {
	go func() {
		http.HandleFunc("/", server.Home)
		http.ListenAndServe(":8080", nil)
	}()

	go func() {
		thermostat.Start()
	}()

	select {}
}
