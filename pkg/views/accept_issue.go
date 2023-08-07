package views

import (
	"html/template"
)

func IssueReqPage() *template.Template {
	temp := template.Must(template.ParseFiles("templates/accept_issue.html", "templates/partials/admin_navbar.html"))
	return temp
}