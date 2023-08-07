package views

import (
	"html/template"
)

func AddAdmin() *template.Template {
	temp := template.Must(template.ParseFiles("templates/add_admin.html", "templates/partials/admin_navbar.html"))
	return temp
}