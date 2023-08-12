package controller

import (
	"mvc/pkg/views"
	"net/http"
)

func AdminServerErrorPage(w http.ResponseWriter, request *http.Request) {
	file := views.FileNames()
	t := views.ViewHomePages(file.AdminServerError)
	t.Execute(w, nil)
}
