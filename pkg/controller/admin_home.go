package controller

import (
	"net/http"
	"mvc/pkg/views"
)

func AdminHome(writer http.ResponseWriter, request *http.Request) {
	t := views.StartAdminPage()
	t.Execute(writer, nil)
}