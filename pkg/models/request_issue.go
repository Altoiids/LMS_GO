package models

import (
	"log"
)

func RequestIssue(username string, bookId int) string {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var userId int
	error := db.QueryRow("SELECT userId FROM user WHERE name=?", username).Scan(&userId)
	if error != nil {
		return "error"
	}

	rows, err := db.Query("SELECT * FROM request WHERE bookId=? AND userId=? ", bookId, userId)
	if err != nil {
		return "error"
	}
	defer rows.Close()

	if rows.Next() {
		return "request already made"
	} else {
		_, err = db.Exec("INSERT INTO request (bookId, userId, status) VALUES (?, ?, ?)", bookId, userId, "issue requested")
		if err != nil {
			return "Database error"
		}
	}
	return ""
}
