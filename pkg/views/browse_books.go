package views

import (
	"html/template"
)

func BrowsePage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/browse_books.html", "templates/partials/user_navbar.html"))
	return temp
}