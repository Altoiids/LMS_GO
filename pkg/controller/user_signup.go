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
	ConfirmPassword := r.FormValue("confirmpassword")

	const AdminID int = 0;

	fmt.Println(Email)
	
	pswd := []byte(Password)
	hashpassword, err := bcrypt.GenerateFromPassword(pswd, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	} 
	Hash := string(hashpassword)
	
	var errorMessage types.ErrorMessage 
	var str string

	str, errorMessage = models.AddUser(AdminID,Name,Email,Hash,Password,ConfirmPassword)

	
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
