package main

import (
	"fmt"
	"net/http"

	"github.com/giovannirossini/scripts/go/web_server/handler"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	http.HandleFunc("/pack", handler.RequestHandler)
	http.HandleFunc("/time", handler.Hour)
	fmt.Println("Listen on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
