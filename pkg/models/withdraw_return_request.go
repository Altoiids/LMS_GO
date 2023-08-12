package models

func WithdrawReturnRequest(username string, bookId int) error {
	var userId int
	db, err := Connection()
	error := db.QueryRow("SELECT userId FROM user WHERE name=?", username).Scan(&userId)
	if error != nil {
		return error
	}

	rows, err := db.Query(`UPDATE request set status = "owned" WHERE bookId = ? and userId=? and status = "return requested";`, bookId, userId)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
