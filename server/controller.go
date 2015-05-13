package server

import (
	"fmt"
	"net/http"
)

// Home displays the web home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
