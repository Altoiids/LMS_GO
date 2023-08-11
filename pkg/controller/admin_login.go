package controller

import (
	"net/http"
	"mvc/pkg/models"
	"mvc/pkg/views"
	)

func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Email := r.FormValue("loginEmail")
	Password := r.FormValue("loginPassword")
    AdminId := 1
	
	string, errorMessage := models.UserLogin(Email,Password,AdminId)
	

		if errorMessage.Message != "no error" {
			t := views.StartAdminPage()
			t.Execute(w, errorMessage)
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:     "jwt",
				Value:    string,
				Path:     "/",
				HttpOnly: true,
			})

			http.Redirect(w, r, "/admin/booksInventory", http.StatusSeeOther)
		}
	}
