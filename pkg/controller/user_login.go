package controller

import (
    "fmt"
	"net/http"
	"mvc/pkg/models"
	"mvc/pkg/views"
)

func LoginUserP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Email := r.FormValue("emaill")
	Password := r.FormValue("passwordl")
	
	fmt.Println(Email,Password)


	str, error := models.UserLogin(Email,Password)

	
	if error != nil {
		t := views.StartPage()
		w.WriteHeader(http.StatusOK)
		t.Execute(w, error)
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