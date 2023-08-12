package controller

import (
	"net/http"
	"mvc/pkg/models"
)

func RemoveAdmin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")

	models.RemoveAdmin(email)
	http.Redirect(w, r, "/admin/viewAdmins", http.StatusSeeOther)
}