package controller

import (
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func AddAdminPage(w http.ResponseWriter, request *http.Request) {
	file := views.FileNames()
	t := views.ViewAdminPages(file.AddAdmin)
	t.Execute(w, nil)
}

func AddAdmin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	const adminId int = 1

	passWord := []byte(password)
	hashpassword, err := bcrypt.GenerateFromPassword(passWord, bcrypt.DefaultCost)
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
	}
	hash := string(hashpassword)

	_, errorMessage := models.AddUser(adminId, name, email, hash, password, confirmPassword)
	if errorMessage.Message != "no error" {
		file := views.FileNames()
		t := views.ViewAdminPages(file.AddAdmin)
		t.Execute(w, errorMessage)
	} else {
		http.Redirect(w, r, "/admin/viewAdmins", http.StatusSeeOther)
	}
}
