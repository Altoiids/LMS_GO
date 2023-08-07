package views

import (
	"html/template"
)

func ListPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/books_inventory.html", "templates/partials/admin_navbar.html"))
	return temp
}