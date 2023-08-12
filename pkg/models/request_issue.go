package models

func RequestIssue(username string, bookId int) error {
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

	rows, err := db.Query("SELECT * FROM request WHERE bookId=? AND userId=? ", bookId, userId)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		return nil
	} else {
		_, err = db.Exec("INSERT INTO request (bookId, userId, status) VALUES (?, ?, ?)", bookId, userId, "issue requested")
		if err != nil {
			return err
		}
	}
	return nil
}
