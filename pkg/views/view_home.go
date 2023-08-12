package views

import (
	"html/template"
)

func ViewHomePages(fileName string) *template.Template {
	temp := template.Must(template.ParseFiles("templates/" + fileName))
	return temp
}