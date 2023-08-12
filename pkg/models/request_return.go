package models

func RequestReturn(username string, bookId int) error {
	db, err := Connection()
	if err != nil {
       return err
	}
	defer db.Close()

	var userId int
	error := db.QueryRow("SELECT userId FROM user WHERE name=?", username).Scan(&userId)
	if error != nil {
		return error
	}

	rows, err := db.Query(`UPDATE request SET status = "return requested" WHERE userId = ? and bookId= ?;`, userId, bookId)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	}
	return nil
}
