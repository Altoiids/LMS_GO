package models

import (
	_ "github.com/go-sql-driver/mysql" 
	"log"
)

func RequestReturn(username string , book_id int) (string){

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
	

	rows, err := db.Query(`UPDATE request SET status = "return requested" WHERE user_id = ? and book_id= ?;`,userID,book_id)
	if err != nil {
		return "error"
	}
	defer rows.Close()

	if rows.Next() {
		return "request already made"
	} 
	return ""
}