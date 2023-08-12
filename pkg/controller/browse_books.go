package controller

import (
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
)

func BrowseBooks(w http.ResponseWriter, request *http.Request) {
	booksList, err := models.FetchBooks()
	if err != nil {
		http.Redirect(w, request, "/client/serverError", http.StatusFound)
		return
	}

	file := views.FileNames()
	t := views.ViewClientPages(file.BrowseBooks)
	w.WriteHeader(http.StatusOK)
	t.Execute(w, booksList)
}
