package views

import (
	"html/template"
)

func UserReturnPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/user_returnrequest.html", "templates/partials/user_navbar.html"))
	return temp
}