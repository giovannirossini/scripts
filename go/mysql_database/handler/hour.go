package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/giovannirossini/curso/advanced/mysql_database/models"
)

// Hour is the handler from request to route /time
func Hour(w http.ResponseWriter, r *http.Request) {
	hour := time.Now().Format("15:04:05")
	body := models.Body{}
	body.Hour = hour
	if err := Model.ExecuteTemplate(w, "index.html", body); err != nil {
		http.Error(w, "There was an error page rendering", http.StatusInternalServerError)
		fmt.Println("[hour] Error in the model execution: ", err.Error())
	}
}
