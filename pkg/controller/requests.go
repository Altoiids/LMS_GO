package controller

import (
	"net/http"
	"mvc/pkg/models"
	"strconv"
)

func IncQty(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	bookid := r.FormValue("bookid")
	bookId, err := strconv.Atoi(bookid)
	
	quantity := r.FormValue("quantity")
	Quantity, err := strconv.Atoi(quantity)
	
	if err != nil {
		return
	}

	models.IncQty(bookId,Quantity)
	http.Redirect(w, r, "/admin/booksinv", http.StatusSeeOther)
}

func DecQty(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	bookid := r.FormValue("bookid")
	bookId, err := strconv.Atoi(bookid)
	
	quantity := r.FormValue("quantity")
	Quantity, err := strconv.Atoi(quantity)
	
	if err != nil {
		return
	}

	models.DecQty(bookId,Quantity)
	http.Redirect(w, r, "/admin/booksinv", http.StatusSeeOther)
}