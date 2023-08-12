package controller

import (
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
)

func LoginUserP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("loginEmail")
	password := r.FormValue("loginPassword")
	adminId := 0

	string, errorMessage := models.UserLogin(email, password, adminId)
	if errorMessage.Message != "no error" {
		file := views.FileNames()
		t := views.ViewHomePages(file.UserHome)
		t.Execute(w, errorMessage)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    string,
			Path:     "/",
			HttpOnly: true,
		})

		http.Redirect(w, r, "/client/profilePage", http.StatusSeeOther)
	}
}
