package models

import (
	_ "github.com/go-sql-driver/mysql" 
	"log"
)

func AddBook(book_name, publisher, isbn string, edition, quantity int) string{

	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM books WHERE book_name=? AND publisher=? AND edition =?", book_name, publisher,edition)
	if err != nil {
		return "There was error"
	}
	defer rows.Close()

	if rows.Next() {
		_, err = db.Exec("UPDATE books SET quantity=quantity+? WHERE book_name=?", quantity, book_name)
		if err != nil {
			return "Database error"
		}
	} else {
		_, err = db.Exec("INSERT INTO books (book_name, publisher, isbn,edition,quantity) VALUES (?, ?, ?,?,?)", book_name, publisher,isbn,edition, quantity)
		if err != nil {
			return "Database error"
		}
	}
	return ""
}