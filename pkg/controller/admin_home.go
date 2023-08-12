package controller

import (
	"mvc/pkg/views"
	"net/http"
)

func AdminHome(w http.ResponseWriter, r *http.Request) {
	file := views.FileNames()
	t := views.ViewHomePages(file.AdminHome)
	t.Execute(w, nil)
}
