package models

import (
	_ "github.com/go-sql-driver/mysql" 
	"log"
)

func RequestIssue(username string , book_id int) (string){

	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var userID int
	error := db.QueryRow("SELECT user_id FROM user WHERE name=?", username).Scan(&userID)
	if error != nil {
		return "error"
	}
	

	rows, err := db.Query("SELECT * FROM request WHERE book_id=? AND user_id=? ", book_id, userID)
	if err != nil {
		return "error"
	}
	defer rows.Close()

	if rows.Next() {
		return "request already made"
	} else {
		_, err = db.Exec("INSERT INTO request (book_id, user_id, status) VALUES (?, ?, ?)", book_id, userID,"issue requested")
		if err != nil {
			return "Database error"
		}
	}
	return ""
}