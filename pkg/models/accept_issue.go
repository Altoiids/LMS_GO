package models

import (
	_ "github.com/go-sql-driver/mysql"
	"mvc/pkg/types"
	"log"
	"fmt"
)

func FetchIssueBooks() ([]types.Book, error) {
	db, err := Connection()
	rows, err := db.Query(`SELECT r.request_id, b.book_id, b.Quantity, b.book_name, u.user_id, u.name FROM request r INNER JOIN user u ON r.user_id = u.user_id INNER JOIN books b ON r.book_id = b.book_id WHERE r.status = "issue requested";
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.RequestID, &book.BookID,&book.Quantity, &book.BookName,&book.UserID, &book.UserName)
		fmt.Println(book.RequestID, book.UserID, book.Quantity,book.UserName,book.BookID, book.BookName)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}


func AcceptIssue(request_id , book_id int) (error){
	
	db, err := Connection()
	if err != nil {
		
		return err
	}
	defer db.Close()
    var available_qty int
	
	error := db.QueryRow("SELECT quantity FROM books WHERE book_id=?", book_id).Scan(&available_qty)
	if error != nil {
		
		return error
	}
    if available_qty == 0{
		
		return error
	} else{
	rows, err := db.Exec(`UPDATE request SET status = "owned" WHERE request_id = ?;` , request_id)
	if err != nil {
		
		return error
		
	}


	rowsAffected, err := rows.RowsAffected()
			if rowsAffected > 0 {
				fmt.Println(book_id)
				_,err := db.Exec(`UPDATE books SET quantity=quantity - 1, issued_qty = issued_qty + 1 WHERE book_id=?`, book_id)
				if err != nil {
				
					return error
				}
			} 
			}
			return error
		
		}




func RejectIssue(request_id int) (string){

	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`DELETE FROM request WHERE request_id = ?`, request_id)
	if err != nil {
		fmt.Println(err)
		return "There was error"
	}
	defer rows.Close()
	
	return ""

}