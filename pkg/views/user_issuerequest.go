package views

import (
	"html/template"
)

func UserIssuePage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/user_issuerequest.html", "templates/partials/user_navbar.html"))
	return temp
}