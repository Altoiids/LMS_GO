package models

import (
	"mvc/pkg/types"
)

func FetchBooks() ([]types.Book, error) {
	db, err := Connection()
	rows, err := db.Query("SELECT bookId, bookName, publisher, isbn,edition,quantity,issuedQuantity FROM books where Quantity >= 0 OR issuedQuantity >=0 ")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookID, &book.BookName, &book.Publisher, &book.ISBN, &book.Edition, &book.Quantity, &book.IssuedQuantity)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
