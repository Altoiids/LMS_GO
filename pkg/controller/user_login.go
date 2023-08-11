package controller

import (
    "fmt"
	"net/http"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"mvc/pkg/types"
)

func LoginUserP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Email := r.FormValue("loginEmail")
	Password := r.FormValue("loginPassword")

	
	fmt.Println(Email,Password)

	var errorMessage types.ErrorMessage 
	var str string
	var AdminId int
    AdminId = 0
	str, errorMessage = models.UserLogin(Email,Password,AdminId)

	
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