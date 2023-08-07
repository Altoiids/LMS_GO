package models

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	
	
)

func WithdrawRR(username string,bookid int) (error) {
	var userID int
	db, err := Connection()
	error := db.QueryRow("SELECT user_id FROM user WHERE name=?", username).Scan(&userID)
	if error != nil {
		return error
	}

   fmt.Println(bookid,userID)

	rows, err := db.Query(`UPDATE request set status = "owned" WHERE book_id = ? and user_id=? and status = "return requested";`,bookid,userID)
	if err != nil {
		fmt.Println("somya recent")
		fmt.Println(err)
		return err
		
	}
	defer rows.Close()

	
	return nil
}