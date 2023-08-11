package models

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"fmt"
)

func IncreaseQuantity( db *sql.DB, book_id,qty int) (string){

	

	rows, err := db.Query("UPDATE books SET Quantity = Quantity + ? WHERE book_id = ?", qty,book_id)
	if err != nil {
		return "There was error"
	}
	defer rows.Close()


	return ""
}

func DecreaseQuantity( book_id,qty int) (string){

	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("UPDATE books SET Quantity = Quantity - ? WHERE book_id = ?", qty,book_id)
	if err != nil {
		return "There was error"
	}
	defer rows.Close()


	return ""
}

func RemoveBook( bookId int) (string){

	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
    fmt.Println("mudit")
	rows, err := db.Query("DELETE FROM books WHERE book_id = ? AND issued_qty = 0",bookId)
	if err != nil {
		return "There was error"
	}
	defer rows.Close()


	return ""
}