package controller

import (
	"net/http"
	"mvc/pkg/models"
	"fmt"
)

func RemoveAdmin(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	email := r.FormValue("email")
	fmt.Println(email)

	models.RemoveAdmin(email)
	http.Redirect(w, r, "/admin/viewAdmins", http.StatusSeeOther)
}