package controller

import (
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	booksList, err := models.FetchBooks()
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}

	file := views.FileNames()
	t := views.ViewAdminPages(file.BooksInventory)
	w.WriteHeader(http.StatusOK)
	t.Execute(w, booksList)
}
