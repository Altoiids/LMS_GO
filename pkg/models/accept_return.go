package models

import (
	_ "github.com/go-sql-driver/mysql"
	"mvc/pkg/types"
	"fmt"
	"log"
)

func FetchReturnBooks() ([]types.Book, error) {
	db, err := Connection()
	rows, err := db.Query(`SELECT r.request_id, b.book_id, b.book_name, u.user_id, u.name FROM request r INNER JOIN user u ON r.user_id = u.user_id INNER JOIN books b ON r.book_id = b.book_id WHERE r.status = "return requested" AND b.Quantity > 0;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.RequestID, &book.BookID, &book.BookName,&book.UserID, &book.UserName)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}


func AcceptReturn(request_id , book_id int) (string){

	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Exec(`DELETE FROM request WHERE request_id = ?` , request_id)
	if err != nil {
		fmt.Println(err)
		return "There was error"
		
	}


	rowsAffected, err := rows.RowsAffected()

	
	if rowsAffected > 0 {
		fmt.Println("somya")
		fmt.Println(book_id)
		_,err := db.Exec(`UPDATE books SET quantity=quantity + 1 WHERE book_id=?`, book_id)
		if err != nil {
			fmt.Println("chawla")
		fmt.Println(err)
		
		}
	} 

	return ""


}


func RejectReturn(request_id int) (string){

	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`UPDATE request SET status = "owned" WHERE request_id = ?;` , request_id)
	if err != nil {
		fmt.Println(err)
		return "There was error"
	}
	defer rows.Close()
	
	return ""

}