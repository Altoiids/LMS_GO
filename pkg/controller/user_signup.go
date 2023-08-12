package controller

import (
	"golang.org/x/crypto/bcrypt"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
)

func AddUserP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	adminId := 0

	passWord := []byte(password)
	hashPassword, err := bcrypt.GenerateFromPassword(passWord, bcrypt.DefaultCost)
	if err != nil {
		http.Redirect(w, r, "/client/serverError", http.StatusFound)
	}
	hash := string(hashPassword)

	str, errorMessage := models.AddUser(adminId, name, email, hash, password, confirmPassword)

	if errorMessage.Message != "no error" {
		file := views.FileNames()
		t := views.ViewHomePages(file.UserHome)
		t.Execute(w, errorMessage)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    str,
			Path:     "/",
			HttpOnly: true,
		})

		http.Redirect(w, r, "/client/profilePage", http.StatusSeeOther)
	}
}
