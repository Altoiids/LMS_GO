package controller

import (
	"encoding/json"
	"net/http"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
)

func AddPage(w http.ResponseWriter, r *http.Request) {
	file := views.FileNames()
	t := views.ViewAdminPages(file.AddBook)
	t.Execute(w, nil)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var body types.Book
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
	}
	
	models.AddBook(body.BookName,body.Publisher,body.ISBN,body.Edition,body.Quantity)
}