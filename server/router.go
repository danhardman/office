package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Router provides the http routing for the application
func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/", GetIndex).Methods("GET")
	r.HandleFunc("/UpdateTemp", UpdateTemp).Methods("POST")

	r.PathPrefix("/css/").Handler(http.StripPrefix("/css", http.FileServer(http.Dir("../static/css"))))
	r.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts", http.FileServer(http.Dir("../static/fonts"))))

	log.Fatal(http.ListenAndServe(":8080", r))
}
