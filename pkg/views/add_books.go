package views

import (
	"html/template"
)

func AddBook() *template.Template {
	temp := template.Must(template.ParseFiles("templates/add_books.html", "templates/partials/admin_navbar.html"))
	return temp
}