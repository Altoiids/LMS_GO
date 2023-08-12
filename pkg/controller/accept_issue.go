package controller

import (
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
	"strconv"
)

func ListIssueRequest(w http.ResponseWriter, r *http.Request) {
	booksList, err := models.FetchIssueBooks()
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}
    file := views.FileNames()
	t := views.ViewAdminPages(file.AcceptIssue)
	w.WriteHeader(http.StatusOK)
	t.Execute(w, booksList)
}

func AcceptIssue(w http.ResponseWriter, r *http.Request) {
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

	error := models.AcceptIssue(requestId, bookId)
	if error != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
	}

	http.Redirect(w, r, "/admin/issueRequests", http.StatusSeeOther)
}

func RejectIssue(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	RequestId := r.FormValue("requestId")
	requestId, err := strconv.Atoi(RequestId)
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}

	error := models.RejectIssue(requestId)
	if error != "" {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/admin/issueRequests", http.StatusSeeOther)
}
