package models

import (
	_ "github.com/go-sql-driver/mysql"
	"mvc/pkg/types"
)

func BrowseBooks() ([]types.Book, error) {
	db, err := Connection()
	rows, err := db.Query("SELECT book_id, book_name, publisher, edition FROM books where Quantity > 0")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookID, &book.BookName, &book.Publisher, &book.Edition)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}