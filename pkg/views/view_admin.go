package views

import (
	"html/template"
)

func ViewAdminPages(fileName string) *template.Template {
	temp := template.Must(template.ParseFiles("templates/" + fileName, "templates/common_components/admin_navbar.html"))
	return temp
}