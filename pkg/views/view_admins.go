package views

import (
	"html/template"
)

func ViewAdmins() *template.Template {
	temp := template.Must(template.ParseFiles("templates/view_admins.html", "templates/partials/admin_navbar.html"))
	return temp
}