package models

import (
	"mvc/pkg/types"
)

func ProfileBooks(username string) ([]types.Book, error) {
	var userId int
	db, err := Connection()
	error := db.QueryRow("SELECT userId FROM user WHERE name=?", username).Scan(&userId)
	if error != nil {
		return nil, error
	}

	rows, err := db.Query("SELECT b.bookId, b.bookName, b.publisher FROM request r JOIN books b ON r.bookId = b.bookId WHERE r.userId =? and r.status = 'owned'", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookID, &book.BookName, &book.Publisher)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}
