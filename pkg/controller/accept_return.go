package controller

import (
	"net/http"
	"strconv"
	"mvc/pkg/models"
	"mvc/pkg/views"
)

func ListReturnRequest(w http.ResponseWriter, r *http.Request) {
	booksList,err := models.FetchReturnBooks()
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}

	file := views.FileNames()
	t := views.ViewAdminPages(file.AcceptReturn)
	w.WriteHeader(http.StatusOK)
	t.Execute(w, booksList)
}

func AcceptReturn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	requestId, err := strconv.Atoi(r.FormValue("requestId"))
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}
	bookId, err := strconv.Atoi(r.FormValue("bookId"))
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}

	error := models.AcceptReturn(requestId,bookId)
	if error != "" {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return 
	}

	http.Redirect(w, r, "/admin/returnRequests", http.StatusSeeOther)
}

func RejectReturn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	RequestId := r.FormValue("requestId")
	requestId, err := strconv.Atoi(RequestId)
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}

	error := models.RejectReturn(requestId)
	if error != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return 
	}
	
	http.Redirect(w, r, "/admin/returnRequests", http.StatusSeeOther)
}