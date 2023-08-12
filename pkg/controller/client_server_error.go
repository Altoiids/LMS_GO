package controller

import (
	"mvc/pkg/views"
	"net/http"
)

func ClientServerErrorPage(w http.ResponseWriter, request *http.Request) {
	file := views.FileNames()
	t := views.ViewHomePages(file.ClientServerError)
	t.Execute(w, nil)
}
