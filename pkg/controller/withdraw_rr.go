package controller

import (
	"net/http"
	"mvc/pkg/models"
	"strings"
	"fmt"
	"strconv"
)

func WithdrawRR(w http.ResponseWriter, r *http.Request) {
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
	BookId := r.FormValue("bookId")
	bookId, err := strconv.Atoi(BookId)


	fmt.Println(username,bookId)
	
	error := models.WithdrawRR(username,bookId)
	if error != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		fmt.Println(error)
		return
	}
	http.Redirect(w, r, "/client/userreturn", http.StatusSeeOther)

}

