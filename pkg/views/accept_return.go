package views

import (
	"html/template"
)

func ReturnReqPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/accept_return.html", "templates/partials/admin_navbar.html"))
	return temp
}