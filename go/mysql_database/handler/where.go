package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/giovannirossini/scripts/go/mysql_database/lib"
	"github.com/giovannirossini/scripts/go/mysql_database/models"
)

// Where is the handler of request to route /where
func Where(w http.ResponseWriter, r *http.Request) {

	where := models.Where{}

	phoneCode, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Do you send a valid number?", http.StatusBadRequest)
		fmt.Println("[where] Error to convert the sended number. Error: ", err.Error())
	}

	sql := "select country, city, phonecode from whereis where phonecode = ?"

	rows, err := lib.Database.Queryx(sql, phoneCode)
	if err != nil {
		http.Error(w, "it was not possible to search this number.", http.StatusBadRequest)
		fmt.Println("[where] it was not possible to execute the query.", sql, " Error: ", err.Error())
		return
	}

	for rows.Next() {
		err = rows.StructScan(&where)
		if err != nil {
			http.Error(w, "it was not possible to search this number.", http.StatusBadRequest)
			fmt.Println("[where] it was not possible to bind the data. Error: ", err.Error())
			return
		}
	}

	if err := ModelWhere.ExecuteTemplate(w, "where.html", where); err != nil {
		http.Error(w, "There was an error page rendering", http.StatusInternalServerError)
		fmt.Println("[hour] Error in the model execution: ", err.Error())
	}
}
