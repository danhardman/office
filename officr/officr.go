package main

import (
	"net/http"

	"github.com/danhardman/officr/app"
	"github.com/danhardman/officr/web"
)

func main() {
	go func() {
		http.HandleFunc("/", web.Home)
		http.ListenAndServe(":8080", nil)
	}()

	app.Start()
}
