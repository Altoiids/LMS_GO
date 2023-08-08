package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"mvc/pkg/models"
	"mvc/pkg/types"
	"mvc/pkg/views"
)

func AddPage(writer http.ResponseWriter, request *http.Request) {
	t := views.AddBook()
	t.Execute(writer, nil)
}

func Add(writer http.ResponseWriter, request *http.Request) {
	var body types.Book
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		fmt.Print("There was an error decoding the request body into the struct")
		fmt.Println(err)
	}
	
	fmt.Printf("Hello Adding the book %s whose publisher is %s, isbn is %s,edition is %d and quantity is %d to the database\n",body.BookName,body.Publisher,body.ISBN,body.Edition,body.Quantity)
	models.AddBook(body.BookName,body.Publisher,body.ISBN,body.Edition,body.Quantity)
}