package controller

import (
	"net/http"
	"time"
)

func LogoutUser(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Expires: time.Now(),
		
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}


func LogoutAdmin(w http.ResponseWriter, r *http.Request) {

	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Expires: time.Now(),
		
	})
	http.Redirect(w, r, "/adminhome", http.StatusSeeOther)
}