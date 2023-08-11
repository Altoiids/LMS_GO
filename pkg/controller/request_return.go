package controller

import (
	"net/http"
	"mvc/pkg/models"
	"strings"
	"strconv"
	"fmt"		
)


func RequestReturn(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	tokenString := strings.TrimSpace(cookie.Value)
	claims, err := models.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	
	username := claims.Username

	r.ParseForm()
	BookId := r.FormValue("book_id")
	bookId, err := strconv.Atoi(BookId)
	fmt.Println(username,bookId)
	models.RequestReturn(username,bookId)
	http.Redirect(w, r, "/client/userreturn", http.StatusSeeOther)
}