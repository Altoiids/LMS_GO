package controller

import (
	"net/http"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"strconv"
	"fmt"
)

func ListIssueRequest(writer http.ResponseWriter, request *http.Request) {
	booksList,err := models.FetchIssueBooks()
	if err != nil {
		http.Error(writer, "Database error", http.StatusInternalServerError)
		return
	}

	t := views.IssueRequestPage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, booksList)
}

func AcceptIssue(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	requestId, err := strconv.Atoi(r.FormValue("requestId"))
	if err != nil {
		return
	}
	bookId, err := strconv.Atoi(r.FormValue("bookId"))
	if err != nil {
		return
	}

	error := models.AcceptIssue(requestId,bookId)
	if error != nil {
	   fmt.Println(error)
	}

	http.Redirect(w, r, "/admin/issueRequests", http.StatusSeeOther)
}

func RejectIssue(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	RequestId := r.FormValue("requestId")
	requestId, err := strconv.Atoi(RequestId)
	if err != nil {
		return
	}
	
	error := models.RejectIssue(requestId)
	if error != "" {
		return 
	}

	http.Redirect(w, r, "/admin/issueRequests", http.StatusSeeOther)
}