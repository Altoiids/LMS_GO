package controller

import (
	"net/http"
	"mvc/pkg/models"
	"strconv"
	"log"
)

func IncreaseQuantity(w http.ResponseWriter, r *http.Request) {
	db, err := models.Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r.ParseForm()
	bookId, err := strconv.Atoi(r.FormValue("bookId"))
	if err != nil {
		return
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		return
	}

	models.IncreaseQuantity(db,bookId,quantity)
	http.Redirect(w, r, "/admin/booksInventory", http.StatusSeeOther)
}

func DecreaseQuantity(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bookId, err := strconv.Atoi(r.FormValue("bookId"))
	if err != nil {
		return
	}
	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		return
	}

	models.DecreaseQuantity(bookId,quantity)
	http.Redirect(w, r, "/admin/booksInventory", http.StatusSeeOther)
}

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	bookId, err := strconv.Atoi(r.FormValue("bookId"))
	if err != nil {
		return
	}

	models.RemoveBook(bookId)
	http.Redirect(w, r, "/admin/booksInventory", http.StatusSeeOther)
}