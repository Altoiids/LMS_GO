package controller

import (
	"mvc/pkg/models"
	"mvc/pkg/views"
	"net/http"
)

func ViewAdmins(w http.ResponseWriter, r *http.Request) {
	db, err := models.Connection()
	rows, err := db.Query("SELECT name, email FROM user where adminId = 1")
	if err != nil {
		return
	}
	defer rows.Close()

	adminsList, err := models.ViewAdmins(db)
	if err != nil {
		http.Redirect(w, r, "/admin/serverError", http.StatusFound)
		return
	}
	file := views.FileNames()
	t := views.ViewAdminPages(file.ViewAdmins)
	w.WriteHeader(http.StatusOK)
	t.Execute(w, adminsList)

}
