package controller

import (
	"net/http"
    "fmt"
	"mvc/pkg/models"
	"mvc/pkg/views"
)

func ViewAdmins(writer http.ResponseWriter, request *http.Request) {
	
	booksList,err := models.ViewAdmins()
	if err != nil {
		http.Error(writer, "Database error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	t := views.ViewAdmins()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, booksList)
	
}



