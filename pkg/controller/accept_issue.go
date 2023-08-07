package controller

import (
	"fmt"
	"net/http"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"strconv"
)

func ListIssueRequest(writer http.ResponseWriter, request *http.Request) {
	booksList,err := models.FetchIssueBooks()
	if err != nil {
		http.Error(writer, "Database error", http.StatusInternalServerError)
		return
	}
	t := views.IssueReqPage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, booksList)
}



func AcceptIssue(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	RequestId := r.FormValue("requestid")
	requestId, err := strconv.Atoi(RequestId)
	
	BookId := r.FormValue("bookid")
	bookId, err := strconv.Atoi(BookId)
	if err != nil {
		return
	}
   fmt.Println(requestId,bookId)
	error := models.AcceptIssue(requestId,bookId)
	if error != "" {
		fmt.Println(error)
		return 
	}

	http.Redirect(w, r, "/admin/issuerequests", http.StatusSeeOther)
}

func RejectIssue(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	RequestId := r.FormValue("requestid")
	requestId, err := strconv.Atoi(RequestId)
	
	
	if err != nil {
		return
	}
	error := models.RejectIssue(requestId)
	if error != "" {
		fmt.Println(error)
		return 
	}

	http.Redirect(w, r, "/admin/issuerequests", http.StatusSeeOther)
}