package models

import (
	_ "github.com/go-sql-driver/mysql"
	"mvc/pkg/types"
	"fmt"
	
)

func UserReturnRequest(username string) ([]types.Book, error) {
	var userID int
	db, err := Connection()
	error := db.QueryRow("SELECT user_id FROM user WHERE name=?", username).Scan(&userID)
	if error != nil {
		return nil, error
	}

	rows, err := db.Query(`SELECT b.book_id, b.book_name, b.publisher FROM request r JOIN books b ON r.book_id = b.book_id WHERE r.user_id = ? and r.status = 'return requested';`,userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookID, &book.BookName, &book.Publisher)
		if err != nil {
			fmt.Println(book.BookID, book.BookName, book.Publisher)
			return nil, err
		}
		books = append(books, book)
		fmt.Println(book.BookID, book.BookName, book.Publisher)
	}
	return books, nil
}