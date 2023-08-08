package controller

import (
	"net/http"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"mvc/pkg/types"
)

func LoginAdminP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	Email := r.FormValue("emaill")
	Password := r.FormValue("passwordl")
	var errorMessage types.ErrorMessage 

	var str string
	var AdminId int
    AdminId = 1
	str, errorMessage = models.UserLogin(Email,Password,AdminId)
	

		if errorMessage.Message != "no error" {
			t := views.StartAdminPage()
			t.Execute(w, errorMessage)
		} else {
			http.SetCookie(w, &http.Cookie{
				Name:     "jwt",
				Value:    str,
				Path:     "/",
				HttpOnly: true,
			})
			http.Redirect(w, r, "/admin/booksinv", http.StatusSeeOther)
		}
	}
