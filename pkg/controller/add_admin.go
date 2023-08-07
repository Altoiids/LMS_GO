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


func AddAdminP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Name := r.FormValue("name")
	Email := r.FormValue("email")
	Password := r.FormValue("password")
	const AdminID int = 1;
	
	
	
	fmt.Println(Email)
    pswd := []byte(Password)
	hashpassword, err := bcrypt.GenerateFromPassword(pswd, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	} 
	hash := string(hashpassword)
	
	
	fmt.Println(err)

	fmt.Printf("Adding the book %s whose publisher is %s, isbn is %s",Name,Email,Password)
	models.AddUser(AdminID,Name,Email,hash)
	http.Redirect(w, r, "/admin/addadmin", http.StatusSeeOther)
}
