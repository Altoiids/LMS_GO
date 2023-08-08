package controller

import (
	"fmt"
	"net/http"
	"mvc/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"mvc/pkg/views"
	"mvc/pkg/types"
)

func AddAdminPage(writer http.ResponseWriter, request *http.Request) {
	t := views.AddAdmin()
	t.Execute(writer, nil)
}

func AddAdminP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Name := r.FormValue("name")
	Email := r.FormValue("email")
	Password := r.FormValue("password")
	ConfirmPassword := r.FormValue("confirmpassword")
	const AdminID int = 1;
    pswd := []byte(Password)
	hashpassword, err := bcrypt.GenerateFromPassword(pswd, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	} 
	Hash := string(hashpassword)
	var errorMessage types.ErrorMessage 
	var returnstr string
	returnstr, errorMessage = models.AddUser(AdminID,Name,Email,Hash,Password,ConfirmPassword)
	fmt.Println(returnstr)
	if errorMessage.Message != "no error" {
		t := views.AddAdmin()
		t.Execute(w, errorMessage)
	} else {
		
		http.Redirect(w, r, "/admin/viewadmins", http.StatusSeeOther)
	}
}
	