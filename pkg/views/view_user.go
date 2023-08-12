package views

import (
	"html/template"
)

func ViewClientPages(fileName string) *template.Template {
	temp := template.Must(template.ParseFiles("templates/" + fileName, "templates/common_components/user_navbar.html"))
	return temp
}