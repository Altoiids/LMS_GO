package models

func WithdrawIssueRequest(username string, bookId int) error {
	var userId int
	db, err := Connection()
	error := db.QueryRow("SELECT userId FROM user WHERE name=?", username).Scan(&userId)
	if error != nil {
		return error
	}
	
	rows, err := db.Query(`DELETE FROM request WHERE bookId = ? and userId=? and status != "owned";`, bookId, userId)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
