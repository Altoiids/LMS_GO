package controller

import (
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func AddPage(w http.ResponseWriter, r *http.Request) {
	file := views.FileNames()
	t := views.ViewAdminPages(file.AddBook)
	t.Execute(w, nil)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var errorMessage types.ErrorMessage

	err2 := r.ParseForm()
	if err2 != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	bookName := r.FormValue("bookName")
	publisher := r.FormValue("publisher")
	isbn := r.FormValue("isbn")
	edition, err := strconv.Atoi(r.FormValue("edition"))
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
	}

	if ValidISBN(isbn) {
		error := models.AddBook(bookName, publisher, isbn, edition, quantity)
		if error != nil {
			http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		}
		http.Redirect(w, r, "/admin/addBook", http.StatusFound)
	} else {
		errorMessage.Message = "ISBN format incorrect"
		file := views.FileNames()
		t := views.ViewAdminPages(file.AddBook)
		t.Execute(w, errorMessage)
	}

}
