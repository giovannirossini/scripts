package handler

import (
	"fmt"
	"net/http"
)

// RequestHandler is a function to handler HTTP requests
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Here is a handler using function in a package.")
}
