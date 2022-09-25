package main

import (
	"fmt"
	"net/http"

	"github.com/giovannirossini/scripts/go/mysql_database/lib"

	"github.com/giovannirossini/scripts/go/mysql_database/handler"
)

func init() {
	fmt.Println("Server up...")
}

func main() {

	err := lib.OpenSQL()
	if err != nil {
		fmt.Println("Server stopping. There was an error on database connection. Error: ", err.Error())
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	})

	http.HandleFunc("/pack", handler.RequestHandler)
	http.HandleFunc("/time", handler.Hour)
	http.HandleFunc("/where/", handler.Where)
	fmt.Println("Listen on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
