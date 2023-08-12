package controller

import (
	"net/http"
	"mvc/pkg/models"
	"strconv"
)

func IncreaseQuantity(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bookId, err := strconv.Atoi(r.FormValue("bookId"))
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}

	models.IncreaseQuantity(bookId,quantity)
	http.Redirect(w, r, "/admin/booksInventory", http.StatusSeeOther)
}

func DecreaseQuantity(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bookId, err := strconv.Atoi(r.FormValue("bookId"))
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}

	models.DecreaseQuantity(bookId,quantity)
	http.Redirect(w, r, "/admin/booksInventory", http.StatusSeeOther)
}

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bookId, err := strconv.Atoi(r.FormValue("bookId"))
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}

	models.RemoveBook(bookId)
	http.Redirect(w, r, "/admin/booksInventory", http.StatusSeeOther)
}