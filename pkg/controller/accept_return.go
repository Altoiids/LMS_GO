package controller

import (
	"net/http"
	"strconv"
	"mvc/pkg/models"
	"mvc/pkg/views"
)

func ListReturnRequest(writer http.ResponseWriter, request *http.Request) {
	booksList,err := models.FetchReturnBooks()
	if err != nil {
		http.Error(writer, "Database error", http.StatusInternalServerError)
		return
	}

	t := views.ReturnReqPage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, booksList)
}

func AcceptReturn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	requestId, err := strconv.Atoi(r.FormValue("requestId"))
	if err != nil {
		return
	}
	bookId, err := strconv.Atoi(r.FormValue("bookId"))
	if err != nil {
		return
	}

	error := models.AcceptReturn(requestId,bookId)
	if error != "" {
		return 
	}

	http.Redirect(w, r, "/admin/returnRequests", http.StatusSeeOther)
}

func RejectReturn(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	RequestId := r.FormValue("requestId")
	requestId, err := strconv.Atoi(RequestId)
	if err != nil {
		return
	}

	error := models.RejectReturn(requestId)
	if error != nil {
		return 
	}
	
	http.Redirect(w, r, "/admin/returnRequests", http.StatusSeeOther)
}