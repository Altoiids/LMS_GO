package views

import (
	"html/template"
)

func StartAdminPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/admin_home.html"))
	return temp
}