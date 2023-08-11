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

func AddBook(writer http.ResponseWriter, request *http.Request) {
	var body types.Book
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
	}
	
	models.AddBook(body.BookName,body.Publisher,body.ISBN,body.Edition,body.Quantity)
}