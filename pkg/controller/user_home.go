package controller

import (
	"net/http"

	"mvc/pkg/views"
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	t := views.StartPage()
	t.Execute(writer, nil)
}