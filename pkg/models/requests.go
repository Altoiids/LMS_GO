package models

func IncreaseQuantity(bookId, quantity int) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("UPDATE books SET quantity = quantity + ? WHERE bookId = ?", quantity, bookId)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func DecreaseQuantity(bookId, quantity int) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("UPDATE books SET quantity = quantity - ? WHERE bookId = ?", quantity, bookId)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func RemoveBook(bookId int) error {
	db, err := Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query("DELETE FROM books WHERE bookId = ? AND issuedQuantity = 0", bookId)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
