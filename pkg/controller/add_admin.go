package controller

import (
	"fmt"
	"net/http"
	"mvc/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"mvc/pkg/views"
)

func AddAdminPage(writer http.ResponseWriter, request *http.Request) {
	t := views.AddAdmin()
	t.Execute(writer, nil)
}

func AddAdmin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	const adminId int = 1;

    passWord := []byte(password)
	hashpassword, err := bcrypt.GenerateFromPassword(passWord, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	} 
	hash := string(hashpassword)

	returnString, errorMessage := models.AddUser(adminId,name,email,hash,password,confirmPassword)
	if errorMessage.Message != "no error" {
		fmt.Println(returnString)
		t := views.AddAdmin()
		t.Execute(w, errorMessage)
	} else {
		http.Redirect(w, r, "/admin/viewAdmins", http.StatusSeeOther)
	}
}
	