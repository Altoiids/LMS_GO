package controller

import (
	"mvc/pkg/views"
	"net/http"
)

func UserLogin(w http.ResponseWriter, request *http.Request) {
	file := views.FileNames()
	t := views.ViewHomePages(file.UserHome)
	t.Execute(w, nil)
}
