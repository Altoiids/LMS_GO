package models

import (
	"log"
)

func AddBook(bookName, publisher, isbn string, edition, quantity int) string {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM books WHERE bookName=? AND publisher=? AND edition =?", bookName, publisher, edition)
	if err != nil {
		return "There was error"
	}
	defer rows.Close()

	if rows.Next() {
		_, err = db.Exec("UPDATE books SET quantity=quantity+? WHERE bookName=? AND publisher=? AND edition =?", quantity, bookName, publisher, edition)
		if err != nil {
			return "Database error"
		}
	} else {
		_, err = db.Exec("INSERT INTO books (bookName, publisher, isbn,edition,quantity) VALUES (?, ?, ?,?,?)", bookName, publisher, isbn, edition, quantity)
		if err != nil {
			return "Database error"
		}
	}
	return ""
}
