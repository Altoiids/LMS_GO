package models

import (
	"log"
)

func IncreaseQuantity(bookId, quantity int) string {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("UPDATE books SET quantity = quantity + ? WHERE bookId = ?", quantity, bookId)
	if err != nil {
		return "There was error"
	}
	defer rows.Close()

	return ""
}

func DecreaseQuantity(bookId, quantity int) string {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("UPDATE books SET quantity = quantity - ? WHERE bookId = ?", quantity, bookId)
	if err != nil {
		return "There was error"
	}
	defer rows.Close()

	return ""
}

func RemoveBook(bookId int) string {
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("DELETE FROM books WHERE bookId = ? AND issuedQuantity = 0", bookId)
	if err != nil {
		return "There was error"
	}
	defer rows.Close()

	return ""
}
