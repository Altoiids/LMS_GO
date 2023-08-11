package controller

import (
	"net/http"
	"mvc/pkg/models"
	"mvc/pkg/views"
	"strings"
	"fmt"
)

func UserReturnRequest(w http.ResponseWriter, r *http.Request) {
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
	
	
	booksList,err := models.UserReturnRequest(username)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	t := views.UserReturnPage()
	w.WriteHeader(http.StatusOK)
	t.Execute(w, booksList)
}

