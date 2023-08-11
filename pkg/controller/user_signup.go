package controller

import (
	"fmt"
	"net/http"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
	"golang.org/x/crypto/bcrypt"
)


func AddUserP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Name := r.FormValue("name")
	Email := r.FormValue("email")
	Password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	const AdminID int = 0;

	fmt.Println(Email)
	
	password := []byte(Password)
	hashpassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	} 
	hash := string(hashpassword)
	
	var errorMessage types.ErrorMessage 
	var str string

	str, errorMessage = models.AddUser(AdminID,Name,Email,hash,Password,confirmPassword)

	
	if errorMessage.Message != "no error" {
		t := views.StartPage()
		t.Execute(w, errorMessage)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    str,
			Path:     "/",
			HttpOnly: true,
		})
		http.Redirect(w, r, "/client/profilepage", http.StatusSeeOther)
	}
}
