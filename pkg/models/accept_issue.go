package models

import (
	_ "github.com/go-sql-driver/mysql"
	"mvc/pkg/types"
	"log"
	"fmt"
)

func FetchIssueBooks() ([]types.Book, error) {
	db, err := Connection()
	rows, err := db.Query(`SELECT r.request_id, b.book_id, b.book_name, u.user_id, u.name FROM request r INNER JOIN user u ON r.user_id = u.user_id INNER JOIN books b ON r.book_id = b.book_id WHERE r.status = "issue requested" AND b.Quantity > 0;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.RequestID, &book.BookID, &book.BookName,&book.UserID, &book.UserName)
		fmt.Println(book.RequestID, book.UserID, book.UserName,book.BookID, book.BookName)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}


func AcceptIssue(request_id , book_id int) (string){
	fmt.Println(request_id,book_id)
	db, err := Connection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Exec(`UPDATE request SET status = "owned" WHERE request_id = ?;` , request_id)
	if err != nil {
		fmt.Println(err)
		return "There was error"
		
	}
	
	rowsAffected, err := rows.RowsAffected()

	
			if rowsAffected > 0 {
				fmt.Println("somya")
				fmt.Println(book_id)
				_,err := db.Exec(`UPDATE books SET quantity=quantity - 1 WHERE book_id=?`, book_id)
				if err != nil {
					fmt.Println("chawla")
				fmt.Println(err)
				
				}
			} 

			return ""
		
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