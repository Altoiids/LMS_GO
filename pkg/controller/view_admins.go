package controller

import (
	"net/http"
    "fmt"
	"mvc/pkg/models"
	"mvc/pkg/views"
)

func ViewAdmins(writer http.ResponseWriter, request *http.Request) {
	db, err := models.Connection()
	rows, err := db.Query("SELECT name, email FROM user where Admin_id = 1")
	if err != nil {
		return
	}
	defer rows.Close()
	adminsList,err := models.ViewAdmins(db)
	if err != nil {
		http.Error(writer, "Database error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	t := views.ViewAdmins()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, adminsList)
	
}



