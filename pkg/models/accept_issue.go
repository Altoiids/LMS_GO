package models

import (
	"mvc/pkg/types"
)

func FetchIssueBooks() ([]types.Book, error) {
	db, err := Connection()
	rows, err := db.Query(`SELECT r.requestId, b.bookId, b.Quantity, b.bookName, u.userId, u.name FROM request r INNER JOIN user u ON r.userId = u.userId INNER JOIN books b ON r.bookId = b.bookId WHERE r.status = "issue requested";`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.RequestID, &book.BookID, &book.Quantity, &book.BookName, &book.UserID, &book.UserName)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func AcceptIssue(requestId, bookId int) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	var availableQuantity int

	error := db.QueryRow("SELECT quantity FROM books WHERE bookId=?", bookId).Scan(&availableQuantity)
	if error != nil {
		return error
	}
	if availableQuantity == 0 {
		return error
	} else {
		rows, err := db.Exec(`UPDATE request SET status = "owned" WHERE requestId = ?;`, requestId)
		if err != nil {
			return error
		}

		rowsAffected, err := rows.RowsAffected()
		if rowsAffected > 0 {
			_, err := db.Exec(`UPDATE books SET quantity=quantity - 1, issuedQuantity = issuedQuantity + 1 WHERE bookId=?`, bookId)
			if err != nil {
				return error
			}
		}
	}
	return error
}

func RejectIssue(requestId int) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query(`DELETE FROM request WHERE requestId = ?`, requestId)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
