package models

import (
	_ "github.com/go-sql-driver/mysql"
	
	"log"
	
)

func IncQty( book_id,qty int) (string){

	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("UPDATE books SET Quantity = Quantity + ? WHERE book_id = ?", qty,book_id)
	if err != nil {
		return "There was error"
	}
	defer rows.Close()


	return ""
}

func DecQty( book_id,qty int) (string){

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